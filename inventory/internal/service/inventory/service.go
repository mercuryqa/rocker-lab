package inventory

import (
	"errors"

	"github.com/mercuryqa/rocket-lab/inventory/internal/repository"
	def "github.com/mercuryqa/rocket-lab/inventory/internal/service"
)

var ErrNotFound = errors.New("part not found")

var _ def.InventoryService = (*InventoryService)(nil)

type InventoryService struct {
	repo repository.InventoryRepository // repo работает с model
}

func NewService(repo repository.InventoryRepository) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}
