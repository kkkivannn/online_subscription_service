package subscriptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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
