package apiv1

import (
	"github.com/mercuryqa/rocket-lab/inventory/internal/service"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

type api struct {
	inventoryV1.UnimplementedInventoryStorageServer
	inventoryService service.InventoryService
}

func NewAPI(inventoryService service.InventoryService) *api {
	return &api{
		inventoryService: inventoryService,
	}
}
