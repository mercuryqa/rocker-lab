package order

import (
	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

func (s *ServiceSuite) TestCancelSuccess() {
	id := "123"
	status := model.Cancelled

	s.OrderRepository.On("CancelOrder", id, status).Return(true)

	ok := s.OrderRepository.CancelOrder(id, status)

	s.Require().True(ok)

	s.OrderRepository.AssertExpectations(s.T())
}

func (s *ServiceSuite) TestCancelNotFound() {
	id := "123"

	status := model.Cancelled

	s.OrderRepository.
		On("CancelOrder", id, status).
		Return(false)

	ok := s.service.CancelOrder(id, status)

	s.Require().False(ok)

	s.OrderRepository.AssertExpectations(s.T())
}
