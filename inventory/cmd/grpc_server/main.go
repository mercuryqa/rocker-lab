package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	apiv1 "github.com/mercuryqa/rocket-lab/inventory/internal/api/inventory/v1"
	repository "github.com/mercuryqa/rocket-lab/inventory/internal/repository/inventory"
	service "github.com/mercuryqa/rocket-lab/inventory/internal/service/inventory"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

const (
	grpcPort = 50055
	envPath  = "../../../deploy/compose/inventory/.env"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("failed to load .env file: %v\n", err)
		return
	}

	dbURI := os.Getenv("MONGO_DB_URI")
	log.Println("Using URI:", dbURI) // для проверки

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Printf("failed to connect to mongo database: %v\n", err)
		return
	}

	defer func() {
		cerr := client.Disconnect(ctx)
		if cerr != nil {
			log.Printf("failed to disconnect mongo database: %v\n", err)
			return
		}
	}()

	// Проверяем соединение с базой данных
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("failed to ping database: %v\n", err)
		return
	}

	// Получаем базу данных
	db := client.Database("inventory")

	// TCP
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}
	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Printf("failed to close listener: %v\n", cerr)
			return
		}
	}()

	// 1️⃣ Создаём репозиторий (слой доступа к данным)
	repo := repository.NewInventoryRepository(db)

	// 2️⃣ Создаём сервис (слой бизнес-логики)
	svc := service.NewService(repo) // InventoryService реализует InventoryStorageServer

	api := apiv1.NewAPI(svc)

	// 3️⃣ Создаём gRPC сервер
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// 4️⃣ Регистрируем сервис
	inventoryV1.RegisterInventoryStorageServer(grpcServer, api)

	go func() {
		log.Printf("🚀 gRPC server listening on %d\n", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			return
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("✅ Server stopped")
}
