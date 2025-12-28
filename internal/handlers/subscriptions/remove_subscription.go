package subscriptions

import (
	"net/http"
	"online_subscription_service/internal/domain/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type DeleteSubscriptionResponse struct {
	Status string `json:"status"`
}

// removeSubscription — HTTP-обработчик для удаления подписки по ID.
//
// Поведение:
//   - Получает ID подписки из URL-параметра
//   - Валидирует и парсит UUID
//   - Вызывает сервисный слой для удаления подписки
//   - Возвращает статус 200 без тела ответа
//
// @Summary     Remove subscription
// @Description Delete subscription from service
// @Tags        subscriptions
// @Accept		json
// @Produce     json
// @Param       id path string true "Subscription ID" format(uuid)
// @Success     200 {string} DeleteSubscriptionResponse
// @Failure     400 {object} models.ErrorResponse "Invalid ID parameter"
// @Failure     500 {object} models.ErrorResponse "Internal server error"
// @Router      /subscriptions/{id} [delete]
func (h *Handlers) removeSubscription(c echo.Context) error {
	param := c.Param("id")

	if param == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "param ID is required"})
	}

	id, err := uuid.Parse(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}

	ctx := c.Request().Context()

	if err := h.subsService.RemoveSubscription(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, DeleteSubscriptionResponse{Status: "Ok"})
}
