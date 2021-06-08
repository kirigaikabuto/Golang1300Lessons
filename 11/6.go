package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	url := "mongodb://localhost:27017"
	databaseName := "db11"
	collectionName := "collection11"
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	db := client.Database(databaseName)
	collection := db.Collection(collectionName)
	p1 := Product{
		Name:  "product1",
		Price: 130,
	}
	result, err := collection.InsertOne(context.TODO(), p1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(result)
}
