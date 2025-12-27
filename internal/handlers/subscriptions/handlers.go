package subscriptions

import (
	"context"
	"net/http"
	"online_subscription_service/internal/domain/models"
	"online_subscription_service/internal/services"

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
	GetPriceWithPeriod(ctx context.Context)
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
	// TODO: создать HTTP-эндпоинт для получения цены с периодом
	h.e.GET("/price", h.getPriceWithPeriod)
}

func (h *Handlers) getPriceWithPeriod(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"code": http.StatusOK,
	})
}
