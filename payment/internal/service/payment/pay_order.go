package payment

import (
	"context"
	"errors"

	"github.com/mercuryqa/rocket-lab/payment/internal/model"
)

var validPaymentMethods = map[model.PaymentMethod]struct{}{
	model.PaymentMethodCard:          {},
	model.PaymentMethodSBP:           {},
	model.PaymentMethodCreditCard:    {},
	model.PaymentMethodInvestorMoney: {},
}

func (s *service) PayOrder(ctx context.Context, info model.PayOrderRequest) (model.PayOrderResponse, error) {
	if _, ok := validPaymentMethods[info.PaymentMethod]; !ok {
		return model.PayOrderResponse{}, errors.New("invalid payment method")
	}

	transactionUuid, err := s.paymentRepository.PayOrder(ctx, info)
	if err != nil {
		return model.PayOrderResponse{}, err
	}
	return transactionUuid, nil
}
