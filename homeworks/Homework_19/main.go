package main

import (
	"GoLandHomeworks/homeworks/Homework_19/product"
	"log"
	"net/http"
)

func main() {
	productsMongo, err := product.NewProductsStore(product.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "product_store",
		CollectionName: "new_products",
	})
	if err != nil {
		log.Fatal(err)
	}
	productHttpEndpoints := product.NewProductHttpEndpoints(productsMongo)
	http.HandleFunc("/", productHttpEndpoints.MainPage)
	http.HandleFunc("/template", productHttpEndpoints.TemplateExample)
	http.HandleFunc("/products", productHttpEndpoints.GetProductsList)
	log.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
