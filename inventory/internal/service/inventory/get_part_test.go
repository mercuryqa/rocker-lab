package inventory

import (
	"errors"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

func (s *ServiceSuite) TestGetPartSuccess() {
	Uuid := "123"

	resp := model.Part{
		UUID:          "123",
		Name:          "name",
		Description:   "description",
		Price:         100.00,
		StockQuantity: 1,
	}

	s.InventoryRepository.
		On("GetPart", s.ctx, Uuid).
		Return(resp, nil)

	part, err := s.service.GetPart(s.ctx, Uuid)

	s.Require().NoError(err)
	s.Require().NotNil(part)
	s.Require().Equal(resp, part)

	s.InventoryRepository.AssertExpectations(s.T())
}

func (s *ServiceSuite) TestGetPartNotFound() {
	id := "missing-order"

	// репозиторий возвращает nil, false
	s.InventoryRepository.
		On("GetPart", s.ctx, id).
		Return(model.Part{}, errors.New("not found"))

	part, err := s.service.GetPart(s.ctx, id)

	s.Require().Error(err)
	s.Require().Equal(model.Part{}, part)

	s.InventoryRepository.AssertExpectations(s.T())
}
