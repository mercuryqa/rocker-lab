package order

import (
	"context"
	"log"

	"github.com/mercuryqa/rocket-lab/order/model"
)

// CreateOrder создает заказ
func (s *service) CreateOrder(ctx context.Context, order *model.GetOrderResponse) error {
	// s.orderRepository.CreateOrder(ctx, order)

	if _, err := s.orderRepository.CreateOrder(ctx, order); err != nil {
		log.Printf("ошибка при создании заказа: %v", err)
		return err
	}

	return nil
}
