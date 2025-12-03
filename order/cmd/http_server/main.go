package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
	"github.com/mercuryqa/rocket-lab/order/internal/api/order/v1"
	grpc "github.com/mercuryqa/rocket-lab/order/internal/client/grpc/db"
	gRPCinventoryV1 "github.com/mercuryqa/rocket-lab/order/internal/client/grpc/inventory/v1"
	gRPCpaymentV1 "github.com/mercuryqa/rocket-lab/order/internal/client/grpc/payment/v1"
	orderRepo "github.com/mercuryqa/rocket-lab/order/internal/repository/order"
	"github.com/mercuryqa/rocket-lab/order/internal/repository/order/db"
	orderService "github.com/mercuryqa/rocket-lab/order/internal/service/order"
	paymentV1 "github.com/mercuryqa/rocket-lab/payment/pkg/proto/payment_v1"
)

const (
	httpPort = "8088"
	// Таймауты для HTTP-сервера
	readHeaderTimeout      = 5 * time.Second
	shutdownTimeout        = 10 * time.Second
	inventoryServerAddress = "localhost:50055"
	paymentServerAddress   = "localhost:50052"
	envPath                = "../../../deploy/compose/order/.env"
)

func main() {
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("⚠️  Не удалось загрузить .env: %v", err)
	}

	dbPool := db.GetDbPool()
	defer dbPool.Close()

	// Инициализируем роутер Chi
	r := chi.NewRouter()

	// Подключаемся к Inventory gRPC-сервису
	invConn := grpc.GRPConn(inventoryServerAddress)
	defer func() {
		if cerr := invConn.Close(); cerr != nil {
			log.Printf("failed to close inventory grpc connection: %v", cerr)
		}
	}()

	// Подключаемся к Payment gRPC-сервису
	payConn := grpc.GRPConn(paymentServerAddress)
	defer func() {
		if cerr := payConn.Close(); cerr != nil {
			log.Printf("failed to close payment grpc connection: %v", cerr)
		}
	}()

	repository := orderRepo.NewOrderRepository(dbPool)
	inventoryClient := gRPCinventoryV1.NewClient(inventoryV1.NewInventoryStorageClient(invConn))
	paymentClient := gRPCpaymentV1.NewClient(paymentV1.NewPaymentV1Client(payConn))

	service := orderService.NewService(repository, inventoryClient, paymentClient)
	handler := apiv1.NewOrderHandler(service)
	handler.RegisterRoutes(r)

	// Запускаем HTTP-сервер
	server := &http.Server{
		Addr:              net.JoinHostPort("localhost", httpPort),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	// Канал который будет ждать закрытия
	stop := make(chan struct{})
	// Запускаем сервер в отдельной горутине
	go func() {
		log.Printf("HTTP-server запущен на порту %s\n", httpPort)
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Ошибка запуска сервера %v\n", err)
		}

		// Закрываем канал
		close(stop)
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Завершение работы сервера...")

	// Создаем контекст с таймаутом для остановки сервера
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Printf("❌ Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("✅ Сервер остановлен")
}
