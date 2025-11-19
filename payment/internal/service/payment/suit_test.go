package payment

import (
	"context"
	"testing"

	"github.com/mercuryqa/rocket-lab/payment/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuit struct {
	suite.Suite

	//nolint:containedctx
	ctx context.Context

	PaymentRepository *mocks.PaymentRepository

	service *service
}

func (s *ServiceSuit) SetupTest() {
	s.ctx = context.Background()

	s.PaymentRepository = mocks.NewPaymentRepository(s.T())

	s.service = NewService(
		s.PaymentRepository,
	)
}

func (s *ServiceSuit) TearDownTest() {}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuit))
}
