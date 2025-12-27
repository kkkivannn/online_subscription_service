package subscriptions

import (
	"context"
	"online_subscription_service/internal/domain/models"
	"online_subscription_service/internal/services"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Service — интерфейс для работы с подписками через HTTP или другие слои.
// Определяет основные операции: создание, чтение, обновление, получение всех подписок,
// получение цены с периодом и удаление подписки.
type Service interface {
	AddSubscription(ctx context.Context, sub models.SubsDTO) (uuid.UUID, error)
	GetSubscription(ctx context.Context, uuid uuid.UUID) (models.Subs, error)
	EditSubscription(ctx context.Context, uuid uuid.UUID, sub models.SubsUpdateDTO) error
	GetAllSubscriptions(ctx context.Context) ([]models.Subs, error)
	GetPriceWithPeriod(ctx context.Context, from, to time.Time, userID uuid.UUID, name string) (int, error)
	RemoveSubscription(ctx context.Context, uuid uuid.UUID) error
}

// Handlers — HTTP-обработчики для работы с подписками.
// Содержит группу маршрутов Echo и ссылку на сервис подписок.
type Handlers struct {
	e           *echo.Group
	subsService *services.SubsService
}

// New — конструктор HTTP-обработчиков.
// Принимает Echo-группу и сервис подписок, возвращает инициализированный Handlers.
func New(
	e *echo.Group,
	subsService *services.SubsService,
) *Handlers {
	return &Handlers{
		e:           e,
		subsService: subsService,
	}
}

// Setup — регистрирует маршруты Echo для работы с подписками.
func (h *Handlers) Setup() {
	h.e.POST("", h.addSubscription)
	h.e.GET("/:id", h.getSubscription)
	h.e.PATCH("/:id", h.editSubscription)
	h.e.DELETE("/:id", h.removeSubscription)
	h.e.GET("", h.getSubscriptions)
	h.e.GET("/price", h.getPriceWithPeriod)
}
