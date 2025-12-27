package subscriptions

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// getSubscription — HTTP-обработчик для получения информации о конкретной подписке.
// - Получает ID подписки из URL-параметра.
// - Валидирует и парсит UUID.
// - Вызывает сервисный слой для получения данных подписки.
// - Возвращает JSON с информацией о подписке или ошибку.
func (h *Handlers) getSubscription(c echo.Context) error {
	param := c.Param("id")

	if param == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "param ID is required",
		})
	}

	id, err := uuid.Parse(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	ctx := c.Request().Context()

	sub, err := h.subsService.GetSubscription(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, sub)
}
