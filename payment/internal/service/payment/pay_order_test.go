package payment

import (
	"context"

	"github.com/mercuryqa/rocket-lab/payment/internal/model"
)

func (s *ServiceSuit) TestPaymentSuccess() {

	ctx := context.Background()

	info := model.PayOrderRequest{
		OrderUuid:     "123",
		UserUuid:      "12",
		PaymentMethod: "PAID",
	}

	resp := model.PayOrderResponse{
		TransactionUuid: "tr-123",
	}

	s.PaymentRepository.
		On(
			"PayOrder",
			ctx,
			info,
		).
		Return(resp, nil)

	transactionUuid, err := s.service.PayOrder(ctx, info)

	s.Require().NoError(err)
	s.Require().NotNil(transactionUuid)
	s.Require().Equal(resp, transactionUuid)

	s.PaymentRepository.AssertExpectations(s.T())
}
