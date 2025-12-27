package services

import (
	"context"
	"online_subscription_service/internal/domain/models"
	"online_subscription_service/internal/storage"

	"github.com/google/uuid"
)

// subsSaver — отвечает за создание подписок.
type subsSaver interface {
	CreateSubscription(ctx context.Context, sub models.SubsDTO) (uuid.UUID, error)
}

// subsProvider — отвечает за чтение и обновление подписок.
type subsProvider interface {
	ReadSubscription(ctx context.Context)
	UpdateSubscription(ctx context.Context)
	ReadAllSubscriptions(ctx context.Context)
	ReadPriceWithPeriod(ctx context.Context)
}

// subsRemover — отвечает за удаление подписок.
type subsRemover interface {
	DeleteSubscriptions(ctx context.Context)
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
	return s.subsSaver.CreateSubscription(ctx, sub)
}

func (s *SubsService) GetSubscription(ctx context.Context) {

}

func (s *SubsService) EditSubscription(ctx context.Context) {

}

func (s *SubsService) GetAllSubscriptions(ctx context.Context) {

}

func (s *SubsService) GetPriceWithPeriod(ctx context.Context) {

}

func (s *SubsService) RemoveSubscription(ctx context.Context) {

}
