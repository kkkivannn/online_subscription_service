package subscriptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	e *echo.Group
}

func New(e *echo.Group) *Handlers {
	return &Handlers{e: e}
}

func (h *Handlers) Setup() {
	h.e.POST("", h.addSubscription)
	h.e.GET("/:id", h.getSubscription)
	h.e.PATCH("/:id", h.editSubscription)
	h.e.DELETE("/:id", h.removeSubscription)
	h.e.GET("", h.getSubscriptions)
	// TODO create http endpoint with period
	h.e.GET("/price", h.getPriceWithPeriod)
}

func (h *Handlers) addSubscription(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *Handlers) getSubscription(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *Handlers) editSubscription(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *Handlers) removeSubscription(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *Handlers) getSubscriptions(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"code": http.StatusOK,
	})
}

func (h *Handlers) getPriceWithPeriod(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"code": http.StatusOK,
	})
}
