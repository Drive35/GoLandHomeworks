package main

import (
	"GoLandHomeworks/homeworks/Homework_14/config"
	"GoLandHomeworks/homeworks/Homework_14/models"
	"fmt"
	"log"
)

func main() {
	cf := config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "online_shop",
		CollectionName: "products",
	}
	collection, err := models.GetConnection(cf)
	if err != nil {
		log.Fatal(err)
	}
	product1 := models.Product{
		Name: "A cruel age",
		Price: 104.27,
		Description: "historical and art book",
	}
	product2 := models.Product{
		Name: "Interstellar",
		Price: 29.99,
		Description: "sci-fi movie",
	}
	product3 := models.Product{
		Name: "Apple iPhone 12 Pro Max",
		Price: 674990.00,
		Description: "the popular smartphone",
	}
	err = models.AddProduct(collection, product1)
	if err != nil {
		log.Fatal(err)
	}
	err = models.AddProduct(collection, product2)
	if err != nil {
		log.Fatal(err)
	}
	err = models.AddProduct(collection, product3)
	if err != nil {
		log.Fatal(err)
	}
	products, err := models.GetAllProducts(collection)
	fmt.Println(products)
}
