package app

import (
	"context"
	"online_subscription_service/internal/config"
	"online_subscription_service/internal/handlers"
	"online_subscription_service/internal/http"
	"online_subscription_service/internal/services"
	"online_subscription_service/internal/storage"
	"online_subscription_service/internal/storage/postgres"

	"github.com/labstack/echo/v4"
)

// App — основной контейнер приложения, оборачивающий HTTP-сервер.
type App struct {
	*http.Server
}

// New — конструктор приложения.
// Выполняет инициализацию всех компонентов приложения.
func New(cfg *config.Config) *App {
	ctx := context.Background()

	e := echo.New()

	srv := http.New(ctx, cfg, e)

	// Подключение к базе данных PostgreSQL.
	db := postgres.New(ctx, cfg)

	// Создание хранилища подписок и сервиса для работы с ними.
	subscriptionsStorage := storage.NewSubsStorage(db)
	subscriptionsService := services.NewSubsService(subscriptionsStorage)

	// Регистрация HTTP-эндпоинтов через Handlers.
	handlers.New(e).SetUpHandlers(subscriptionsService)

	return &App{srv}
}
