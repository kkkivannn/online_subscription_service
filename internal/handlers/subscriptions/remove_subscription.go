package subscriptions

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// removeSubscription — HTTP-обработчик для удаления подписки по ID.
// - Получает ID подписки из URL-параметра.
// - Валидирует и парсит UUID.
// - Вызывает сервисный слой для удаления подписки.
// - Возвращает статус 200 при успешном удалении или соответствующую ошибку.
func (h *Handlers) removeSubscription(c echo.Context) error {
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

	if err := h.subsService.RemoveSubscription(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
