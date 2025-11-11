package inventory

import (
	"errors"
	"sync"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	def "github.com/mercuryqa/rocket-lab/inventory/internal/repository"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

var ErrNotFound = errors.New("part not found")

var _ def.InventoryRepository = (*InventoryRepository)(nil)

// InventoryStorage представляет потокобезопасное хранилище данных о заказах
type InventoryRepository struct {
	inventoryV1.UnimplementedInventoryStorageServer

	mu        sync.RWMutex
	inventory map[string]*model.GetPartResponse
}

func NewInventoryRepository() *InventoryRepository {
	s := &InventoryRepository{
		inventory: make(map[string]*model.GetPartResponse),
	}
	GenerateSampleData(s)
	return s
}
