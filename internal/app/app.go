package app

import (
	"context"
	"online_subscription_service/internal/config"
	"online_subscription_service/internal/handlers"
	"online_subscription_service/internal/http"

	"github.com/labstack/echo/v4"
)

type App struct {
	*http.Server
}

func New(cfg *config.Config) *App {
	ctx := context.Background()

	e := echo.New()

	srv := http.New(ctx, cfg, e)

	h := handlers.New(e)
	h.SetUpHandlers()

	return &App{srv}
}
