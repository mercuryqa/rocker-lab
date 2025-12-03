package gen

import (
	"context"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertOrderIDs(db *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1️⃣ Очищаем таблицу перед вставкой
	_, err := db.Exec(ctx, "DELETE FROM items")
	if err != nil {
		return err
	}

	ids := make([]string, 0, 17)

	for i := 1; i <= 17; i++ {
		ids = append(ids, strconv.Itoa(i))
	}

	// Создаём builder
	builder := sq.Insert("items").
		Columns("part_uuid"). // только один столбец
		PlaceholderFormat(sq.Dollar)

	// Добавляем каждое значение через Values
	for _, id := range ids {
		builder = builder.Values(id)
	}

	// Генерируем SQL и аргументы
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	// Выполняем вставку
	_, err = db.Exec(ctx, query, args...)
	return err
}
