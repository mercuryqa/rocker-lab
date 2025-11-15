package order

import (
	"log"

	"github.com/mercuryqa/rocket-lab/order/internal/converter"
	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

func (r *OrderRepository) GetOrder(id string) (*model.Order, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, ok := r.orders[id]
	if !ok {
		return nil, false
	}

	log.Print(order)

	return converter.RepoModelToModel(order), true
}
