package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	apiv1 "github.com/mercuryqa/rocket-lab/inventory/internal/api/inventory/v1"
	repository "github.com/mercuryqa/rocket-lab/inventory/internal/repository/inventory"
	service "github.com/mercuryqa/rocket-lab/inventory/internal/service/inventory"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

const (
	grpcPort = 50055
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Printf("failed to close listener: %v\n", cerr)
		}
	}()

	// 1️⃣ Создаём репозиторий (слой доступа к данным)
	repo := repository.NewInventoryRepository()

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
