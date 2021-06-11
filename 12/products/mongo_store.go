package products

import (
	"context"
	"github.com/google/uuid"
	"github.com/kirigaikabuto/Golang1300Lessons/12/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//структура в которой хранятся элементы которые позволяют совершать какие либо операции в бд
type ProductStore struct {
	collection *mongo.Collection
}

//метод который создает (сущность в которой хранятся элементы которые позволяют совершать какие либо операции в бд)
//а так же подключается к базе данных

func NewProductStore(config config.MongoConfig) (*ProductStore, error) {
	clientOptions := options.Client().ApplyURI(config.Url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(config.Database)
	collection := db.Collection(config.Collection)
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

func (p *ProductStore) List() ([]Product, error) {
	filter := bson.D{}
	products := []Product{}
	cursor, err := p.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
