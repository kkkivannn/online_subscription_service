package http

import (
	"context"
	"fmt"
	"online_subscription_service/internal/config"

	"github.com/labstack/echo/v4"
)

// Server — структура, описывающая HTTP-сервер.
// Включает контекст для graceful shutdown, порт, хост и экземпляр Echo.
type Server struct {
	ctx context.Context
	cfg *config.Config
	e   *echo.Echo
}

func New(
	ctx context.Context,
	cfg *config.Config,
	e *echo.Echo,
) *Server {
	return &Server{
		ctx: ctx,
		cfg: cfg,
		e:   e,
	}
}

// MustRun — запускает HTTP-сервер и завершает приложение с фатальной ошибкой, если запуск невозможен.
// Формирует строку адреса (хост:порт) и запускает Echo.
func (s *Server) MustRun() {
	host := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	s.e.Logger.Fatal(s.e.Start(host))
}

// Stop — корректно завершает работу сервера (graceful shutdown) с использованием контекста.
// Вызывается при завершении приложения или получении сигнала завершения.
func (s *Server) Stop(ctx context.Context) {
	if err := s.e.Shutdown(ctx); err != nil {
		s.e.Logger.Error("Can't stop server")
	}
}
