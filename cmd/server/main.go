package main

import (
	"context"
	"log/slog"
	"online_subscription_service/internal/app"
	"online_subscription_service/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// main является точкой входа в приложение.
//  1. Загружает конфигурацию.
//  2. Инициализирует и запускает приложение.
//  3. Ожидает сигнала завершения (SIGINT или SIGTERM).
//  4. Корректно завершает работу приложения с таймаутом в 10 секунд.
func main() {
	cfg := config.MustLoad()

	app := app.New(cfg)

	go app.MustRun()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	slog.Info("shutdown initiated...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Останавливаем приложение.
	app.Stop(shutdownCtx)
	slog.Info("server stopped...")
}
