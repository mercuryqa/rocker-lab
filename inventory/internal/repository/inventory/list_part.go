package inventory

import (
	"context"
	"log"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	"github.com/mercuryqa/rocket-lab/inventory/internal/repository/converter"
)

func (r *InventoryRepository) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	repoFilter := converter.PartsFilterToRepo(filter)
	cursor, err := r.collection.Find(ctx, repoFilter)
	if err != nil {
		log.Printf("Ошибка получения данных по фильтру: фильтр - %v, ошибка - %v\n", filter, err)
	}
	defer func() {
		if cerr := cursor.Close(ctx); cerr != nil {
			log.Printf("Ошибка при закрытии курсора: %v\n", cerr)
		}
	}()

	var parts []model.Part
	err = cursor.All(ctx, &parts)
	if err != nil {
		log.Printf("Ошибка декодирования parts: %v\n", err)
	}
	// list := []model.Part{}
	// for _, part := range parts {
	//	list[part.UUID] = part
	// }
	if len(parts) == 0 {
		return []model.Part{}, model.ErrPartListEmpty
	}
	return parts, nil
}
