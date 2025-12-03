package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/mercuryqa/rocket-lab/order/internal/migrator"
)

func GetDbPool() *pgxpool.Pool {
	ctx := context.Background()

	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		log.Fatal("DB_URI is not set in environment")
	}

	pool, err := pgxpool.New(ctx, dbURI)
	if err != nil {
		log.Fatalf("Ошибка создания пула подключений: %v", err)
	}

	// Проверим соединение
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}

	// Прогоняем миграции
	migrationDir := os.Getenv("MIGRATION_DIR")
	if migrationDir == "" {
		migrationDir = "migrations"
	}

	migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*pool.Config().ConnConfig.Copy()), migrationDir)

	if err := migratorRunner.Up(); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	return pool
}
