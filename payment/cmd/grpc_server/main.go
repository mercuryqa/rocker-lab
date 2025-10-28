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

	paymentV1API "github.com/mercuryqa/rocket-lab/payment/internal/api/payment/v1"
	paymentRepository "github.com/mercuryqa/rocket-lab/payment/internal/repository/payment"
	paymentService "github.com/mercuryqa/rocket-lab/payment/internal/service/payment"
	paymentV1 "github.com/mercuryqa/rocket-lab/payment/pkg/proto/payment_v1"
)

const (
	grpcPort = 50052
)

//type paymentService struct {
//	paymentV1.UnimplementedPaymentV1Server
//
//	mu       sync.RWMutex
//	payments map[string]*paymentV1.PaymentMethod
//}

// PayOrder
//func (s *paymentService) PayOrder(_ context.Context, req *paymentV1.PayOrderRequest) (*paymentV1.PayOrderResponse, error) {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//
//	// Генерируем новый UUID для транзакции
//	transactionUUID := uuid.NewString()
//
//	log.Printf("Создана транзакция %s для заказа %s", transactionUUID, req.OrderUuid)
//
//	// Возвращаем ответ клиенту
//	return &paymentV1.PayOrderResponse{
//		TransactionUuid: transactionUUID,
//	}, nil
//
//}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}
	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Printf("failed to close listener: %v\n", cerr)
		}
	}()

	// Создаем gRPC сервер
	s := grpc.NewServer()

	repo := paymentRepository.NewRepository()
	service := paymentService.NewService(repo)
	api := paymentV1API.NewAPI(service)

	// Регистрируем наш сервис
	//service := &paymentService{
	//	payments: make(map[string]*paymentV1.PaymentMethod),
	//}

	paymentV1.RegisterPaymentV1Server(s, api)

	// Включаем рефлексию для отладки
	reflection.Register(s)

	go func() {
		log.Printf("🚀 gRPC server listening on %d\n", grpcPort)
		err = s.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down gRPC server...")
	s.GracefulStop()
	log.Println("✅ Server stopped")
}
