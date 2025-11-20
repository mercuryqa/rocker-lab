package order

import (
	"sync"

	def "github.com/mercuryqa/rocket-lab/order/internal/repository"
	repoModel "github.com/mercuryqa/rocket-lab/order/internal/repository/model"
)

var _ def.OrderRepository = (*OrderRepository)(nil)

// Представляет потокобезопасное хранилище данных о заказах
type OrderRepository struct {
	mu     sync.Mutex
	orders map[string]*repoModel.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]*repoModel.Order),
	}
}
