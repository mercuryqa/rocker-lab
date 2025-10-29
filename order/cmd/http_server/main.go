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

	"github.com/mercuryqa/rocket-lab/order/internal/api/order/v1"
	orderRepo "github.com/mercuryqa/rocket-lab/order/internal/repository/order"
	orderService "github.com/mercuryqa/rocket-lab/order/internal/service/order"
)

const (
	httpPort = "8088"
	// Таймауты для HTTP-сервера
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

func main() {
	// Инициализируем роутер Chi
	r := chi.NewRouter()

	repository := orderRepo.NewOrderRepository()
	service := orderService.NewService(repository)
	handler := apiV1.NewOrderHandler(service)
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
		log.Printf("HTTP-server запущен не порту %s\n", httpPort)
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
