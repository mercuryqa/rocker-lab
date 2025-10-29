// internal/api/order_handler.go

package apiV1

import (
	"strings"

	orderService "github.com/mercuryqa/rocket-lab/order/internal/service"
	paymentV1 "github.com/mercuryqa/rocket-lab/payment/pkg/proto/payment_v1"
)

type OrderHandler struct {
	service orderService.OrderService
}

func NewOrderHandler(service orderService.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
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
