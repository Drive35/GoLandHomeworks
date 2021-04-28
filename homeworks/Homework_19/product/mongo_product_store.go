package product

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type ProductsStore interface {
	AddProduct(product Product) error
	GetAllProducts() ([]Product, error)
}
type productsStore struct {
	collection *mongo.Collection
}

func NewProductsStore(cf MongoConfig) (ProductsStore, error) {
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
	return &productsStore{collection: collection}, nil
}

func (p *productsStore) AddProduct(product Product) error {
	_, err := p.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	return nil
}

func (p *productsStore) GetAllProducts() ([]Product, error) {
	products := []Product{}
	cursor, err := p.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
