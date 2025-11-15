package inventory

import (
	"github.com/mercuryqa/rocket-lab/inventory/internal/repository"
	def "github.com/mercuryqa/rocket-lab/inventory/internal/service"
)

var _ def.InventoryService = (*InventoryService)(nil)

type InventoryService struct {
	repo repository.InventoryRepository // repo работает с model
}

func NewService(repo repository.InventoryRepository) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}
