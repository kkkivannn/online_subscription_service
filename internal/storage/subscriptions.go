package storage

import (
	"context"
	"fmt"
	"online_subscription_service/internal/domain/models"

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

func (s *SubsStorage) ReadSubscription(ctx context.Context) {

}

func (s *SubsStorage) UpdateSubscription(ctx context.Context) {

}

func (s *SubsStorage) ReadAllSubscriptions(ctx context.Context) {

}

func (s *SubsStorage) ReadPriceWithPeriod(ctx context.Context) {

}

func (s *SubsStorage) DeleteSubscriptions(ctx context.Context) {

}
