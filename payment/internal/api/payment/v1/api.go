package api_v1

import (
	"github.com/mercuryqa/rocket-lab/payment/internal/service"
	paymentV1 "github.com/mercuryqa/rocket-lab/payment/pkg/proto/payment_v1"
)

type api struct {
	paymentV1.UnimplementedPaymentV1Server

	paymentService service.PaymentService
}

func NewAPI(paymentService service.PaymentService) *api {
	return &api{
		paymentService: paymentService,
	}
}
