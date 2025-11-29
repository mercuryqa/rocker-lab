package inventory

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

func (r *InventoryRepository) GetPart(ctx context.Context, uuid string) (model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var part model.Part
	err := r.collection.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&part)
	if err != nil {
		log.Printf("Ошибка получения part из коллекции: %v\n", err)
		return model.Part{}, model.ErrNotFound
	}
	return part, nil
}
