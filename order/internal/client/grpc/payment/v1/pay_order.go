package v1

import (
	"context"

	generatedPaymentV1 "github.com/mercuryqa/rocket-lab/payment/pkg/proto/payment_v1"
)

func (c *client) PayOrder(ctx context.Context, orderUUID, userUUID string, paymentMethod generatedPaymentV1.PaymentMethod) (string, error) {
	res, err := c.generatedClient.PayOrder(ctx, &generatedPaymentV1.PayOrderRequest{
		OrderUuid:     orderUUID,
		UserUuid:      userUUID,
		PaymentMethod: paymentMethod,
	})

	if err != nil {
		return "", err
	}

	return res.TransactionUuid, nil
}
