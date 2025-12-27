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
	GetSubscription(ctx context.Context)
	EditSubscription(ctx context.Context)
	GetAllSubscriptions(ctx context.Context)
	GetPriceWithPeriod(ctx context.Context)
	RemoveSubscription(ctx context.Context)
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

// addSubscription — HTTP-обработчик для создания новой подписки.
// Принимает данные запроса, конвертирует их в SubsDTO и вызывает сервис.
// Возвращает JSON с UUID новой подписки или ошибку.
func (h *Handlers) addSubscription(c echo.Context) error {
	r := new(models.SubRequest)
	ctx := c.Request().Context()
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	id, err := h.subsService.AddSubscription(ctx, *r.ToSubsDTO())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"id": id.String(),
	})
}

func (h *Handlers) getSubscription(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *Handlers) editSubscription(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *Handlers) removeSubscription(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *Handlers) getSubscriptions(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"code": http.StatusOK,
	})
}

func (h *Handlers) getPriceWithPeriod(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"code": http.StatusOK,
	})
}
