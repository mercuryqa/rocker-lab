package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

func (s *ServiceSuite) TestInventorySuccess() {
	ctx := context.Background()

	Uuid := "123"

	resp := model.Part{
		UUID:          "123",
		Name:          "name",
		Description:   "description",
		Price:         100.00,
		StockQuantity: 1,
	}

	s.InventoryRepository.
		On("GetPart", ctx, Uuid).
		Return(resp, nil)

	part, err := s.service.GetPart(ctx, Uuid)

	s.Require().NoError(err)
	s.Require().NotNil(part)
	s.Require().Equal(resp, part)

	s.InventoryRepository.AssertExpectations(s.T())
}
