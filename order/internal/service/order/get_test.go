package order

import (
	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

func (s *ServiceSuite) TestGetOrderSuccess() {
	id := "order-123"

	expectedOrder := &model.Order{
		OrderUuid:  id,
		UserUuid:   "user-1",
		PartUuids:  []string{"p1", "p2"},
		TotalPrice: 150.0,
	}

	// мок ожидаем, что репозиторий вернёт именно этот order
	s.OrderRepository.
		On("GetOrder", id).
		Return(expectedOrder, true)

	order, ok := s.service.GetOrder(id)

	s.Require().True(ok)
	s.Require().NotNil(order)
	s.Require().Equal(expectedOrder, order)

	s.OrderRepository.AssertExpectations(s.T())
}

func (s *ServiceSuite) TestGetOrderNotFound() {
	id := "missing-order"

	// репозиторий возвращает nil, false
	s.OrderRepository.
		On("GetOrder", id).
		Return((*model.Order)(nil), false)

	order, ok := s.service.GetOrder(id)

	s.Require().False(ok)
	s.Require().Nil(order)

	s.OrderRepository.AssertExpectations(s.T())
}
