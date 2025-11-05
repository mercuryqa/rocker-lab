package order

import (
	"sync"

	def "github.com/mercuryqa/rocket-lab/order/internal/repository"
	"github.com/mercuryqa/rocket-lab/order/model"
)

var _ def.OrderRepository = (*OrderRepository)(nil)

// Представляет потокобезопасное хранилище данных о заказах
type OrderRepository struct {
	mu     sync.Mutex
	orders map[string]*model.GetOrderResponse
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]*model.GetOrderResponse),
	}
}
