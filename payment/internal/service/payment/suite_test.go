package payment

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/mercuryqa/rocket-lab/payment/internal/repository/mocks"
)

type ServiceSuite struct {
	suite.Suite

	//nolint:containedctx
	ctx context.Context

	PaymentRepository *mocks.PaymentRepository

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.PaymentRepository = mocks.NewPaymentRepository(s.T())

	s.service = NewService(
		s.PaymentRepository,
	)
}

func (s *ServiceSuite) TearDownTest() {}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
