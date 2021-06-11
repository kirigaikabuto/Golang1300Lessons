package main

import (
	"fmt"
	"github.com/kirigaikabuto/Golang1300Lessons/12/config"
	"github.com/kirigaikabuto/Golang1300Lessons/12/products"
	"log"
)

func main() {
	productStore, err := products.NewProductStore(config.MongoConfig{
		Url:        "mongodb://localhost:27017",
		Database:   "lesson12",
		Collection: "products",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	product, err := productStore.GetById("2c613b4c-79dd-4024-8ad8-3d7b2b2f6566")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(product)
	//	p := &products.Product{
	//		Name:  "product2",
	//		Price: 123,
	//	}
	//	newProduct, err := productStore.Create(p)
	//	if err != nil {
	//		log.Fatal(err)
	//		return
	//	}
	//	fmt.Println(newProduct)
	//	productsObjects, err := productStore.List()
	//	if err != nil {
	//		log.Fatal(err)
	//		return
	//	}
	//	fmt.Println(productsObjects)
}
