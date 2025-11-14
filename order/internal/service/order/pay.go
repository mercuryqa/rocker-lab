package order

import (
	"context"
	"log"
	"strings"
	"time"

	paymentV1 "github.com/mercuryqa/rocket-lab/payment/pkg/proto/payment_v1"
)

func (s *service) PayOrder(id, paymentMethod string) string {
	order, ok := s.GetOrder(id)
	if !ok {
		log.Print("Order not found")
		return ""
	}
	if order.Status == "PAID" {
		log.Print("Order already PAID")
		return ""
	}
	if order.Status == "CANCELED" {
		log.Print("Order canceled - can't be paid")
		return ""
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	paymentMethodName := toPaymentMethod(paymentMethod)
	transactionUuid, err := s.paymentClient.PayOrder(ctx, order.OrderUuid, order.UserUuid, paymentMethodName)
	if err != nil {
		log.Printf("grpc payment service call failed: %v", err)
		return ""
	}

	s.orderRepository.UpdateOrder(id, paymentMethod, transactionUuid)
	s.orderRepository.PayOrder(id, "PAID")

	return transactionUuid
}

func toPaymentMethod(s string) paymentV1.PaymentMethod {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "CARD":
		return paymentV1.PaymentMethod_CARD
	case "SBP":
		return paymentV1.PaymentMethod_SBP
	case "CREDIT_CARD":
		return paymentV1.PaymentMethod_CREDIT_CARD
	case "INVESTOR_MONEY":
		return paymentV1.PaymentMethod_INVESTOR_MONEY
	default:
		return paymentV1.PaymentMethod_UNKNOWN
	}
}
