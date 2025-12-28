package postgres

import (
	"context"
	"fmt"
	"online_subscription_service/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v5/pgxpool"
)

// New создает и возвращает новое подключение к пулу PostgreSQL.
// Принимает контекст выполнения и указатель на конфигурацию приложения.
func New(context context.Context, config *config.Config) *pgxpool.Pool {
	dbHost := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.DB.Username, // Имя пользователя
		config.DB.Password, // Пароль
		config.DB.DBHost,   // Хост базы данных
		config.DB.DBPort,   // Порт базы данных
		config.DB.DBName,   // Название базы данных
		config.DB.SSLMode,
	)
	var conn *pgxpool.Pool
	var err error

	count := 10

	for range count {
		conn, err = pgxpool.New(context, dbHost)
		if err == nil {
			err = conn.Ping(context)
		}

		if err == nil {
			break
		}
	}
	m, err := migrate.New("file://./migrations", dbHost)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	return conn
}
