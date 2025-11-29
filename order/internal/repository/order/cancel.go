package order

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *OrderRepository) CancelOrder(ctx context.Context, id, status string) bool {
	log.Printf("id: %v\n status: %v\n", id, status)

	builderUpdate := sq.Update("orders").
		Set("status", "CANCEL").
		Where(sq.Eq{"order_uuid": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Printf("failed to build query")
		return false
	}

	_, err = r.poolDb.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("Ошибка insert в таблицу orders: %v\n", err)
		return false
	}

	return true
}
