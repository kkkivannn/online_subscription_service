package subscriptions

import (
	"net/http"
	"online_subscription_service/internal/domain/models"

	"github.com/labstack/echo/v4"
)

// addSubscription — HTTP-обработчик для создания новой подписки.
// Принимает данные запроса, конвертирует их в SubsDTO и вызывает сервис.
// Возвращает JSON с UUID новой подписки или ошибку.
func (h *Handlers) addSubscription(c echo.Context) error {
	r := new(models.AddSubRequest)
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
