package subscriptions

import (
	"net/http"
	"online_subscription_service/internal/domain/models"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GetPriceWithPeriodResponse struct {
	Price int `json:"price"`
}

// getPriceWithPeriod — HTTP-обработчик для получения стоимости подписки за указанный период.
//
// Параметры запроса:
//   - from: начало периода (формат "YYYY-MM-DD").
//   - to: конец периода (формат "YYYY-MM-DD").
//   - user_id: UUID пользователя.
//   - service_name: название услуги.
//
// Валидирует входные параметры, парсит даты и UUID, вызывает сервисный слой.
// Возвращает JSON с рассчитанной ценой или ошибку.
//
// GetPriceWithPeriod godoc
// @Summary Get price
// @Description Get price with period from service
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param       from query string true "Start date" format(date) example(2025-01-01)
// @Param       to query string true "End date" format(date) example(2025-01-31)
// @Param       user_id query string true "User ID" format(uuid) example(550e8400-e29b-41d4-a716-446655440000)
// @Param       service_name query string true "Service name" example("premium")
// @Success     200 {object} GetPriceWithPeriodResponse
// @Failure     400 {object} models.ErrorResponse "Invalid request parameters"
// @Failure     500 {object} models.ErrorResponse "Internal server error"
// @Router      /subscriptions/price [get]
func (h *Handlers) getPriceWithPeriod(c echo.Context) error {
	fromStr := c.QueryParam("from")
	if fromStr == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "param From is required"})
	}

	from, err := time.Parse("2006-01-02", fromStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid from date"})
	}

	toStr := c.QueryParam("to")
	if toStr == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "param To is required"})
	}

	to, err := time.Parse("2006-01-02", toStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid to date"})
	}

	id := c.QueryParam("user_id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "param user_id is required"})
	}

	userID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
	}

	serviceName := c.QueryParam("service_name")
	if serviceName == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "param service_name is required"})
	}

	ctx := c.Request().Context()

	price, err := h.subsService.GetPriceWithPeriod(ctx, from, to, userID, serviceName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, GetPriceWithPeriodResponse{Price: price})
}
