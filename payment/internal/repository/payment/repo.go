package payment

import (
	"sync"

	def "github.com/mercuryqa/rocket-lab/payment/internal/repository"
	repoModel "github.com/mercuryqa/rocket-lab/payment/internal/repository/model"
)

var _ def.PaymentRepository = (*repository)(nil)

type repository struct {
	mu   sync.RWMutex
	data map[string]repoModel.PayOrderResponse
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]repoModel.PayOrderResponse),
	}
}
