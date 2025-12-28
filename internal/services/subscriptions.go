package services

import (
	"context"
	"errors"
	"log/slog"
	"online_subscription_service/internal/domain/models"
	"online_subscription_service/internal/storage"
	"time"

	"github.com/google/uuid"
)

// subsSaver — отвечает за создание подписок.
type subsSaver interface {
	CreateSubscription(ctx context.Context, sub models.SubsDTO) (uuid.UUID, error)
}

// subsProvider — отвечает за чтение и обновление подписок.
type subsProvider interface {
	ReadSubscription(ctx context.Context, uuid uuid.UUID) (models.SubsDTO, error)
	UpdateSubscription(ctx context.Context, uuid uuid.UUID, sub models.SubsUpdateDTO) error
	ReadAllSubscriptions(ctx context.Context) ([]models.SubsDTO, error)
	ReadPriceWithPeriod(ctx context.Context, from, to time.Time, userID uuid.UUID, name string) (int, error)
}

// subsRemover — отвечает за удаление подписок.
type subsRemover interface {
	DeleteSubscriptions(ctx context.Context, uuid uuid.UUID) error
}

// SubsService — сервисный слой для работы с подписками.
// Объединяет возможности создания, чтения/обновления и удаления подписок через соответствующие интерфейсы.
type SubsService struct {
	subsSaver    subsSaver
	subsProvider subsProvider
	subsRemover  subsRemover
}

// NewSubsService — конструктор сервиса подписок.
// Принимает хранилище подписок и возвращает инициализированный экземпляр SubsService,
// реализующий все операции через соответствующие интерфейсы.
func NewSubsService(subsStorage *storage.SubsStorage) *SubsService {
	return &SubsService{
		subsSaver:    subsStorage,
		subsProvider: subsStorage,
		subsRemover:  subsStorage,
	}
}

// AddSubscription — добавляет новую подписку через интерфейс subsSaver.
// Возвращает UUID созданной подписки и ошибку, если она произошла.
func (s *SubsService) AddSubscription(ctx context.Context, sub models.SubsDTO) (uuid.UUID, error) {
	slog.Info("start adding subscription")
	uuid, err := s.subsSaver.CreateSubscription(ctx, sub)
	if err != nil {
		slog.Error(err.Error())
		return uuid, errors.New("error add new subscription")
	}
	return uuid, nil
}

// GetSubscription — возвращает информацию о конкретной подписке по UUID.
// Вызывает метод subsProvider.ReadSubscription и конвертирует результат в модель Subs.
func (s *SubsService) GetSubscription(ctx context.Context, uuid uuid.UUID) (models.Subs, error) {
	slog.Info("start getting subscription")
	sub, err := s.subsProvider.ReadSubscription(ctx, uuid)
	if err != nil {
		slog.Error(err.Error())
		return models.Subs{}, errors.New("error get subscription")
	}

	return sub.ToSubs(), nil
}

// EditSubscription — обновляет данные существующей подписки.
// Вызывает метод subsProvider.UpdateSubscription с переданным UUID и DTO обновления.
func (s *SubsService) EditSubscription(ctx context.Context, uuid uuid.UUID, sub models.SubsUpdateDTO) error {
	slog.Info("start editting subscription")
	if err := s.subsProvider.UpdateSubscription(ctx, uuid, sub); err != nil {
		slog.Error(err.Error())
		return errors.New("error edit subscription")
	}
	return nil
}

// GetAllSubscriptions — возвращает список всех подписок.
// Читает данные через subsProvider.ReadAllSubscriptions и конвертирует каждую запись в модель Subs.
func (s *SubsService) GetAllSubscriptions(ctx context.Context) ([]models.Subs, error) {
	slog.Info("start getting all subscriptions")
	var subs []models.Subs

	slice, err := s.subsProvider.ReadAllSubscriptions(ctx)
	if err != nil {
		slog.Error(err.Error())
		return subs, errors.New("error geting all subscriptions")
	}

	for _, v := range slice {
		subs = append(subs, v.ToSubs())
	}

	return subs, nil
}

// GetPriceWithPeriod — возвращает стоимость подписки за указанный период для конкретного пользователя и услуги.
// Параметры: начало и конец периода, UUID пользователя, название услуги.
// Вызывает subsProvider.ReadPriceWithPeriod для вычисления цены.
func (s *SubsService) GetPriceWithPeriod(ctx context.Context, from, to time.Time, userID uuid.UUID, name string) (int, error) {
	slog.Info("start getting price with period")
	price, err := s.subsProvider.ReadPriceWithPeriod(ctx, from, to, userID, name)
	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("error getting price with period")
	}
	return price, nil
}

// RemoveSubscription — удаляет подписку по UUID.
// Вызывает метод subsRemover.DeleteSubscriptions для удаления записи.
func (s *SubsService) RemoveSubscription(ctx context.Context, uuid uuid.UUID) error {
	slog.Info("start deleting subscription")
	if err := s.subsRemover.DeleteSubscriptions(ctx, uuid); err != nil {
		slog.Error(err.Error())
		return errors.New("error deleting subscription")
	}
	return nil
}
