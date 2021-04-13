package models

import (
	"GoLandHomeworks/homeworks/Homework_14/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type Product struct {
	Name        string
	Price       float64
	Description string
}

func GetConnection(cf config.MongoConfig) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + cf.Host + ":" + cf.Port)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(cf.Database)
	err = db.CreateCollection(context.TODO(), cf.CollectionName)
	if err != nil && !strings.Contains(err.Error(), "NamespaceExists") {
		return nil, err
	}
	collection := db.Collection(cf.CollectionName)
	return collection, nil
}
func AddProduct(collection *mongo.Collection, product Product) error {
	_, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	return nil
}

func GetAllProducts(collection *mongo.Collection) ([]Product, error) {
	products := []Product{}
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
