package handlers

import (
	"online_subscription_service/internal/handlers/subscriptions"
	"online_subscription_service/internal/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handlers — основной контейнер для HTTP-обработчиков приложения.
// Содержит экземпляр Echo для регистрации маршрутов и middleware.
type Handlers struct {
	e *echo.Echo
}

// New — конструктор Handlers.
// Принимает экземпляр Echo и возвращает инициализированный Handlers.
func New(e *echo.Echo) *Handlers {
	return &Handlers{e: e}
}

// SetUpHandlers — настраивает все HTTP-эндпоинты и middleware приложения.
func (h *Handlers) SetUpHandlers(subscriptionsService *services.SubsService) {
	// Восстанавливает приложение после паники и логирует ошибки
	h.e.Use(middleware.Recover())

	// Включает логгирование всех HTTP-запросов
	h.e.Use(middleware.RequestLogger())

	// Группа всех API-эндпоинтов с префиксом /api/v1
	api := h.e.Group("/api/v1")

	// Группа эндпоинтов для подписок (/api/v1/subscriptions)
	subs := api.Group("/subscriptions")
	subscriptions.New(subs, subscriptionsService).Setup()
}
