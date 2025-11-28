package v1

import (
	"context"

	generatedInventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
	clientConverter "github.com/mercuryqa/rocket-lab/order/internal/client/converter"
	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

func (c *client) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	parts, err := c.generatedClient.ListParts(ctx, &generatedInventoryV1.GetListPartRequest{
		Filter: clientConverter.PartsFilterToProto(filter),
	})
	if err != nil {
		return nil, err
	}

	return clientConverter.PartListToModel(parts.Parts), nil
}
