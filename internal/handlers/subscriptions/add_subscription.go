package subscriptions

import (
	"net/http"
	"online_subscription_service/internal/domain/models"

	"github.com/labstack/echo/v4"
)

type AddSubscriptionResponse struct {
	ID string `json:"id"`
}

// addSubscription — HTTP-обработчик для создания новой подписки.
// Принимает данные запроса, конвертирует их в SubsDTO и вызывает сервис.
// Возвращает JSON с UUID новой подписки или ошибку.
//
// AddSubscription godoc
// @Summary Add subscription
// @Description Add subscription to service
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param        request body models.AddSubRequest true "Subscription data"
// @Success      201 {object} AddSubscriptionResponse
// @Failure      400 {object} models.ErrorResponse
// @Failure      500 {object} models.ErrorResponse
// @Router       /subscriptions [post]
func (h *Handlers) addSubscription(c echo.Context) error {
	r := new(models.AddSubRequest)
	ctx := c.Request().Context()

	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
	}

	id, err := h.subsService.AddSubscription(ctx, *r.ToSubsDTO())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusCreated, AddSubscriptionResponse{ID: id.String()})
}
