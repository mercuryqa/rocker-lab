package v1

import (
	def "github.com/mercuryqa/rocket-lab/order/internal/client/grpc"
	generatedPaymentV1 "github.com/mercuryqa/rocket-lab/payment/pkg/proto/payment_v1"
)

var _ def.PaymentClient = (*client)(nil)

// Implement interface
type client struct {
	generatedClient generatedPaymentV1.PaymentV1Client
}

func NewClient(generatedClient generatedPaymentV1.PaymentV1Client) *client {
	return &client{
		generatedClient: generatedClient,
	}
}
