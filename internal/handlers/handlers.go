package handlers

import (
	"online_subscription_service/internal/handlers/subscriptions"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	e *echo.Echo
}

func New(e *echo.Echo) *Handlers {
	return &Handlers{e: e}
}

func (h *Handlers) SetUpHandlers() {
	// Восстанавливает приложение после паники и логирует ошибки
	h.e.Use(middleware.Recover())

	// Включает логгирование всех HTTP-запросов
	h.e.Use(middleware.RequestLogger())

	// Группа всех API-эндпоинтов с префиксом /api/v1
	api := h.e.Group("/api/v1")

	// Группа эндпоинтов для подписок (/api/v1/subscriptions)
	subs := api.Group("/subscriptions")
	subscriptions.New(subs).Setup()
}
