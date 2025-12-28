package subscriptions

import (
	"net/http"
	"online_subscription_service/internal/domain/models"

	"github.com/labstack/echo/v4"
)

// getSubscriptions — HTTP-обработчик для получения списка подписок.
//
// Поведение:
//   - Вызывает сервисный слой для получения всех подписок
//   - Возвращает массив подписок
//
// @Summary     Get subscriptions
// @Description Get subscriptions from service
// @Tags        subscriptions
// @Accept		json
// @Produce     json
// @Success     200 {array} models.Subs
// @Failure     500 {object} models.ErrorResponse "Internal server error"
// @Router      /subscriptions [get]
func (h *Handlers) getSubscriptions(c echo.Context) error {
	ctx := c.Request().Context()

	subs, err := h.subsService.GetAllSubscriptions(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "param ID is required"})
	}

	return c.JSON(http.StatusOK, subs)
}
