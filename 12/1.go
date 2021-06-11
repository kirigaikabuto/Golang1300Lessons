package main

import (
	"fmt"
	"github.com/kirigaikabuto/Golang1300Lessons/12/products"
	"log"
)

func main() {
	productStore, err := products.NewProductStore(
		"mongodb://localhost:27017",
		"lesson12",
		"products",
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	p := &products.Product{
		Name:  "product2",
		Price: 123,
	}
	newProduct, err := productStore.Create(p)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(newProduct)
	productsObjects, err := productStore.List()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(productsObjects)
}
