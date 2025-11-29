package order

import (
	"context"
	"log"
)

func (s *service) CheckItems(ctx context.Context, ids []string) bool {
	ok := s.orderRepository.CheckItems(ctx, ids)
	if !ok {
		log.Println("No ITEMS")
		return false
	}

	return true
}
