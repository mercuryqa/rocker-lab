package inventory

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	def "github.com/mercuryqa/rocket-lab/inventory/internal/repository"
	repoModel "github.com/mercuryqa/rocket-lab/inventory/internal/repository/model"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

var _ def.InventoryRepository = (*InventoryRepository)(nil)

// InventoryStorage представляет потокобезопасное хранилище данных о заказах
type InventoryRepository struct {
	inventoryV1.UnimplementedInventoryStorageServer
	mu sync.RWMutex

	collection *mongo.Collection
}

func NewInventoryRepository(db *mongo.Database) *InventoryRepository {
	collection := db.Collection("parts")

	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "uuid", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		panic(err)
	}

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Ошибка получения данных из MongoDB: %b\n", err)
	}
	if cerr := cursor.Close(ctx); cerr != nil {
		log.Printf("Ошибка закрытия курсора: %v\n", cerr)
	}

	var val []repoModel.Part
	err = cursor.All(ctx, &val)
	if err != nil {
		log.Printf("Ошибка записи данных из MongoDB в слайс: %v\n", err)
	}
	if val == nil {
		_, err = collection.InsertMany(ctx, GetCollectionParts())
		if err != nil {
			log.Printf("Ошибка заполнения MongoDB: %v\n", err)
		}
	}

	return &InventoryRepository{
		collection: collection,
	}
}
