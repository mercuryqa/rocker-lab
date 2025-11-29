package order

import (
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	gen "github.com/mercuryqa/rocket-lab/order/internal/gen"
	"github.com/mercuryqa/rocket-lab/order/internal/migrator"
	def "github.com/mercuryqa/rocket-lab/order/internal/repository"
)

var _ def.OrderRepository = (*OrderRepository)(nil)

// Представляет потокобезопасное хранилище данных о заказах
type OrderRepository struct {
	poolDb *pgxpool.Pool
}

func NewOrderRepository(poolDb *pgxpool.Pool) *OrderRepository {
	migrationsDir := os.Getenv("MIGRATION_DIR")
	if migrationsDir == "" {
		migrationsDir = "order/migrations"
	}

	// Прогоняем миграции
	pgxConfig := poolDb.Config().ConnConfig.Copy()
	sqlDB := stdlib.OpenDB(*pgxConfig)

	migratorRunner := migrator.NewMigrator(sqlDB, migrationsDir)
	if err := migratorRunner.Up(); err != nil {
		log.Printf("Ошибка миграции базы данных: %v\n", err)
		panic(err)
	}

	log.Println("✅ Миграции успешно применены")

	err := gen.InsertOrderIDs(poolDb)
	if err != nil {
		log.Printf("failed to fill table items: %v\n", err)
		return nil
	}

	return &OrderRepository{
		poolDb: poolDb,
	}
}
