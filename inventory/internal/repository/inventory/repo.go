package inventory

import (
	"sync"

	def "github.com/mercuryqa/rocket-lab/inventory/internal/repository"
	repoModel "github.com/mercuryqa/rocket-lab/inventory/internal/repository/model"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

var _ def.InventoryRepository = (*InventoryRepository)(nil)

// InventoryStorage представляет потокобезопасное хранилище данных о заказах
type InventoryRepository struct {
	inventoryV1.UnimplementedInventoryStorageServer

	mu        sync.RWMutex
	inventory map[string]*repoModel.GetPartResponse
}

func NewInventoryRepository() *InventoryRepository {
	s := &InventoryRepository{
		inventory: make(map[string]*repoModel.GetPartResponse),
	}
	generateSampleData(s)
	return s
}
