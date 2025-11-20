package v1

import (
	generatedInventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
	def "github.com/mercuryqa/rocket-lab/order/internal/client/grpc"
)

var _ def.InventoryClient = (*client)(nil)

type client struct {
	generatedClient generatedInventoryV1.InventoryStorageClient
}

func NewClient(generatedClient generatedInventoryV1.InventoryStorageClient) *client {
	return &client{
		generatedClient: generatedClient,
	}
}
