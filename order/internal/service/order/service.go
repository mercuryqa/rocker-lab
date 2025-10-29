package order

import (
	"github.com/mercuryqa/rocket-lab/order/internal/repository"
	def "github.com/mercuryqa/rocket-lab/order/internal/service"
	"github.com/mercuryqa/rocket-lab/order/model"
)

var _ def.OrderService = (*service)(nil)

type service struct {
	orderRepository repository.OrderRepository
}

func NewService(orderRepository repository.OrderRepository) *service {
	return &service{
		orderRepository: orderRepository,
	}
}

// CreateOrder создает заказ
func (s *service) CreateOrder(order *model.GetOrderResponse) {
	s.orderRepository.CreateOrder(order)
}

func (s *service) GetOrder(id string) (*model.GetOrderResponse, bool) {
	order, ok := s.orderRepository.GetOrder(id)
	if !ok {
		return nil, false
	}
	return order, true
}

func (s *service) PayOrder(id, status string) bool {
	s.orderRepository.PayOrder(id, status)
	return true
}

func (s *service) CanсelOrder(id, status string) bool {
	s.orderRepository.CanсelOrder(id, status)
	return true
}

func (s *service) UpdateOrder(id, paymentMethod, transactionUuid string) bool {
	s.orderRepository.UpdateOrder(id, paymentMethod, transactionUuid)
	return true
}
