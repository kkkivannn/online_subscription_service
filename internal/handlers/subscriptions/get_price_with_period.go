package subscriptions

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// getPriceWithPeriod — HTTP-обработчик для получения стоимости подписки за указанный период.
// Параметры запроса:
// - from: начало периода (формат "YYYY-MM-DD").
// - to: конец периода (формат "YYYY-MM-DD").
// - user_id: UUID пользователя.
// - service_name: название услуги.
// Валидирует входные параметры, парсит даты и UUID, вызывает сервисный слой.
// Возвращает JSON с рассчитанной ценой или ошибку.
func (h *Handlers) getPriceWithPeriod(c echo.Context) error {
	fromStr := c.QueryParam("from")
	if fromStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "from is required",
		})
	}

	from, err := time.Parse("2006-01-02", fromStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "invalid from date"})
	}

	toStr := c.QueryParam("to")
	if toStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "to is required",
		})
	}

	to, err := time.Parse("2006-01-02", toStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "invalid to date"})
	}

	id := c.QueryParam("user_id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "user_id is required",
		})
	}

	userID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	serviceName := c.QueryParam("service_name")
	if serviceName == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "service_name is required",
		})
	}

	ctx := c.Request().Context()

	price, err := h.subsService.GetPriceWithPeriod(ctx, from, to, userID, serviceName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"price": price,
	})
}
