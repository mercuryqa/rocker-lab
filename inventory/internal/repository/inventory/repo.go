package inventory

import (
	"errors"
	"sync"

	def "github.com/mercuryqa/rocket-lab/inventory/internal/repository"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

var ErrNotFound = errors.New("part not found")

var _ def.InventoryRepository = (*InventoryRepository)(nil)

// InventoryStorage представляет потокобезопасное хранилище данных о заказах
type InventoryRepository struct {
	inventoryV1.UnimplementedInventoryStorageServer

	mu        sync.RWMutex
	inventory map[string]*inventoryV1.GetPartResponse
}

func NewInventoryRepository() *InventoryRepository {
	s := &InventoryRepository{
		inventory: make(map[string]*inventoryV1.GetPartResponse),
	}
	GenerateSampleData(s)
	return s
}
