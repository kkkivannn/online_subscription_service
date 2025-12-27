package storage

import (
	"context"
	"errors"
	"fmt"
	"online_subscription_service/internal/domain/models"
	"online_subscription_service/internal/lib/storage"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SubsStorage — хранилище для работы с подписками.
// Содержит подключение к базе данных и реализует методы
// для создания и управления записями подписок.
type SubsStorage struct {
	db *pgxpool.Pool
}

// NewSubsStorage — конструктор хранилища подписок.
// Принимает пул подключений к базе данных и
// возвращает инициализированный экземпляр SubsStorage.
func NewSubsStorage(db *pgxpool.Pool) *SubsStorage {
	return &SubsStorage{
		db: db,
	}
}

// CreateSubscription — создает новую подписку в таблице services.
// На вход принимает структуру SubsDTO с данными подписки
// (название, цена, идентификатор пользователя, дата начала и окончания).
// Возвращает UUID созданной подписки.
// Если при выполнении запроса произошла ошибка, возвращает её наружу.
func (s *SubsStorage) CreateSubscription(ctx context.Context, sub models.SubsDTO) (uuid.UUID, error) {
	var ID uuid.UUID

	query := "insert into services (name, price, user_id, start_date, end_date) values ($1, $2, $3, $4, $5) returning id"

	err := s.db.QueryRow(ctx, query, sub.Name, sub.Price, sub.UserID.String(), sub.StartDate, sub.EndDate).Scan(&ID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to insert sub: %w", err)
	}

	return ID, nil
}

// ReadSubscription — читает подписку по UUID из базы данных.
// Возвращает DTO подписки или ошибку, если запись не найдена или произошла ошибка при запросе.
func (s *SubsStorage) ReadSubscription(ctx context.Context, uuid uuid.UUID) (models.SubsDTO, error) {
	var sub models.SubsDTO

	query := "select id, name, price, user_id, start_date, end_date from services where id=$1"

	err := s.db.QueryRow(ctx, query, uuid).Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate)
	if err != nil {
		return sub, fmt.Errorf("failed to select sub: %w", err)
	}
	return sub, nil
}

// UpdateSubscription — обновляет существующую подписку по UUID.
// Использует BuildUpdateQuery для генерации SQL-запроса и аргументов.
// Возвращает ошибку, если не удалось обновить запись или аргументы пусты.
func (s *SubsStorage) UpdateSubscription(ctx context.Context, uuid uuid.UUID, sub models.SubsUpdateDTO) error {
	query, args := storage.BuildUpdateQuery(uuid, sub)
	if args == nil {
		return errors.New("empty args for update")
	}

	data, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update service: %w", err)
	}

	if data.RowsAffected() == 0 {
		return fmt.Errorf("failed to update service, 0 rows affected")
	}

	return nil
}

// ReadAllSubscriptions — возвращает все подписки из базы данных.
// Конвертирует каждую запись в DTO и возвращает срез подписок.
func (s *SubsStorage) ReadAllSubscriptions(ctx context.Context) ([]models.SubsDTO, error) {
	var subs []models.SubsDTO

	query := "select * from services"

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return subs, fmt.Errorf("failed to select subs: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sub models.SubsDTO
		err := rows.Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate)
		if err != nil {
			return subs, fmt.Errorf("failed to scan sub: %w", err)
		}
		subs = append(subs, sub)
	}

	return subs, nil
}

// ReadPriceWithPeriod — вычисляет суммарную стоимость подписки за указанный период для конкретного пользователя и услуги.
// Возвращает 0, если подписок в периоде нет.
func (s *SubsStorage) ReadPriceWithPeriod(ctx context.Context, from, to time.Time, userID uuid.UUID, name string) (int, error) {
	var price int
	query := "select coalesce(sum(price),0) from services where user_id = $1 and name=$2 and start_date < $3 ::date + interval '1 day'  and (end_date is null or end_date >= $4)"

	err := s.db.QueryRow(ctx, query, userID, name, to, from).Scan(&price)
	if err != nil {
		return 0, fmt.Errorf("failed to read price: %w", err)
	}

	return price, nil
}

// DeleteSubscriptions — удаляет подписку по UUID из базы данных.
// Возвращает ошибку, если запись не найдена или произошла ошибка при удалении.
func (s *SubsStorage) DeleteSubscriptions(ctx context.Context, uuid uuid.UUID) error {
	query := "delete from services where id=$1"

	data, err := s.db.Exec(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}

	if data.RowsAffected() == 0 {
		return fmt.Errorf("failed to delete service, 0 rows affected")
	}

	return nil
}
