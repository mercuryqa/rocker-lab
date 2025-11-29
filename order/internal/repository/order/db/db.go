package db

//
// import (
//	"context"
//	"log"
//	"os"
//
//	"github.com/jackc/pgx/v5"
//	"github.com/jackc/pgx/v5/pgxpool"
//	"github.com/jackc/pgx/v5/stdlib"
//	"github.com/joho/godotenv"
//
//	"github.com/mercuryqa/rocket-lab/order/internal/migrator"
// )
//
//// const envPath = "/Users/pkirchanov/data/github/microservices/.env"
//
// func GetDbConn() *pgx.Conn {
//	ctx := context.Background()
//
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Printf("failed to load .env file: %v\n", err)
//		return nil
//	}
//
//	dbURI := os.Getenv("DB_URI")
//
//	// Создаем соединение с БД
//	conn, err := pgx.Connect(ctx, dbURI)
//	if err != nil {
//		log.Printf("failed to connect to database: %v\n", err)
//		return nil
//	}
//
//	defer func() {
//		cerr := conn.Close(ctx)
//		if cerr != nil {
//			log.Printf("failed to close connection: %v\n", cerr)
//			return
//		}
//	}()
//
//	// Проверяем, что соединение с базой установлено
//	err = conn.Ping(ctx)
//	if err != nil {
//		log.Printf("База данных не доступна: %v\n", err)
//		return nil
//	}
//
//	// Инициализируем мигратор
//	migrationDir := os.Getenv("MIGRATION_DIR")
//
//	// Преобразование из pgx в sql.DB
//	migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*conn.Config().Copy()), migrationDir)
//
//	err = migratorRunner.Up()
//	if err != nil {
//		log.Printf("Ошибка миграции базы данных: %v\n", err)
//		return nil
//	}
//
//	return conn
// }
//
// func GetDbPool() *pgxpool.Pool {
// 	ctx := context.Background()
//
//	dbURI := os.Getenv("DB_URI")
//
//	pool, err := pgxpool.New(ctx, dbURI)
//	if err != nil {
//		log.Printf("Ошибка подключения к б/д: %v\n", err)
//	}
//
//	return pool
// }
