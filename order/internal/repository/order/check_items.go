package order

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *OrderRepository) CheckItems(ctx context.Context, ids []string) bool {
	builderSelect := sq.Select("COUNT(*)").
		From("items").
		Where(sq.Eq{"part_uuid": ids}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to build query")
		return false
	}

	var count int
	err = r.poolDb.QueryRow(ctx, query, args...).Scan(&count)
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return false
	}

	if count < len(ids) {
		return false
	}
	return true
}
