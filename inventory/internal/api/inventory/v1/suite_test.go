package apiv1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/mercuryqa/rocket-lab/inventory/internal/service/mocks"
)

type APISuite struct {
	suite.Suite

	//nolint:containedctx
	ctx context.Context

	inventoryService *mocks.InventoryService

	api *api
}

func (a *APISuite) SetupTest() {
	a.ctx = context.Background()

	a.inventoryService = mocks.NewInventoryService(a.T())

	a.api = NewAPI(
		a.inventoryService,
	)
}

func (a *APISuite) TearDownTest() {}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(APISuite))
}
