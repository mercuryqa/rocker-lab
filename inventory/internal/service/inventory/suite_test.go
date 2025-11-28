package inventory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/mercuryqa/rocket-lab/inventory/internal/repository/mocks"
)

type ServiceSuite struct {
	suite.Suite

	//nolint:containedctx
	ctx context.Context

	InventoryRepository *mocks.InventoryRepository

	service *InventoryService
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.InventoryRepository = mocks.NewInventoryRepository(s.T())

	s.service = NewService(
		s.InventoryRepository,
	)
}

func (s *ServiceSuite) TearDownTest() {}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
