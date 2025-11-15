package order

import (
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

	return converter.RepoModelToModel(order), true
}
