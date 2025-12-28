package subscriptions

import (
	"net/http"
	"online_subscription_service/internal/domain/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// getSubscription — HTTP-обработчик для получения информации о конкретной подписке.
//
//   - Получает ID подписки из URL-параметра.
//   - Валидирует и парсит UUID.
//   - Вызывает сервисный слой для получения данных подписки.
//   - Возвращает JSON с информацией о подписке или ошибку.
//
// GetSubscription godoc
// @Summary Get subscription
// @Description Get subscription from service
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param        id path string true "Subscription ID" format(uuid)
// @Success      200 {object} models.Subs
// @Failure      400 {object} models.ErrorResponse
// @Failure      500 {object} models.ErrorResponse
// @Router       /subscriptions/{id} [get]
func (h *Handlers) getSubscription(c echo.Context) error {
	param := c.Param("id")

	if param == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "param ID is required"})
	}

	id, err := uuid.Parse(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}

	ctx := c.Request().Context()

	sub, err := h.subsService.GetSubscription(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, sub)
}
