package products

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//структура в которой хранятся элементы которые позволяют совершать какие либо операции в бд
type ProductStore struct {
	collection *mongo.Collection
}

//метод который создает (сущность в которой хранятся элементы которые позволяют совершать какие либо операции в бд)
//а так же подключается к базе данных

func NewProductStore(url, databaseName, collectionName string) (*ProductStore, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(databaseName)
	collection := db.Collection(collectionName)
	return &ProductStore{collection: collection}, nil
}

func (p *ProductStore) Create(product *Product) (*Product, error) {
	id := uuid.New()
	product.Id = id.String()
	_, err := p.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
