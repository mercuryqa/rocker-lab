package order

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	client "github.com/mercuryqa/rocket-lab/order/internal/client/grpc/mocks"
	"github.com/mercuryqa/rocket-lab/order/internal/repository/mocks"
)

type ServiceSuite struct {
	suite.Suite

	//nolint:containedctx
	ctx context.Context

	OrderRepository *mocks.OrderRepository
	InventoryClient *client.InventoryClient
	PaymentClient   *client.PaymentClient

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.OrderRepository = mocks.NewOrderRepository(s.T())
	s.InventoryClient = client.NewInventoryClient(s.T())
	s.PaymentClient = client.NewPaymentClient(s.T())

	s.service = NewService(
		s.OrderRepository,
		s.InventoryClient,
		s.PaymentClient,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
