package order

import (
	"sync"

	def "github.com/mercuryqa/rocket-lab/order/internal/repository"
	"github.com/mercuryqa/rocket-lab/order/model"
)

var _ def.OrderRepository = (*OrderService)(nil)

// представляет потокобезопасное хранилище данных о заказах
type OrderService struct {
	mu     sync.Mutex
	orders map[string]*model.GetOrderResponse
}

func NewOrderStorage() *OrderService {
	return &OrderService{
		orders: make(map[string]*model.GetOrderResponse),
	}
}
