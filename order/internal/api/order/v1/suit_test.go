package apiv1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/mercuryqa/rocket-lab/order/internal/service/mocks"
)

type APISuite struct {
	suite.Suite

	//nolint:containedctx
	ctx context.Context

	orderService *mocks.OrderService

	api *OrderHandler
}

func (a *APISuite) SetupTest() {
	a.ctx = context.Background()

	a.orderService = mocks.NewOrderService(a.T())

	a.api = NewOrderHandler(
		a.orderService,
	)
}

func (a *APISuite) TearDownTest() {}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(APISuite))
}
