package subscriptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// getSubscriptions — HTTP-обработчик для получения всех подписок.
// - Вызывает сервисный слой для получения списка всех подписок.
// - Возвращает JSON с массивом подписок или ошибку.
func (h *Handlers) getSubscriptions(c echo.Context) error {
	ctx := c.Request().Context()

	subs, err := h.subsService.GetAllSubscriptions(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, subs)
}
