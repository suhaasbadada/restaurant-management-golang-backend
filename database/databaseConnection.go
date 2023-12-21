package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDbURI := "mongodb://localhost:27017"
	fmt.Println(MongoDbURI)

	// Use mongo.Connect directly with a context and client options
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDbURI))
	handleError(err)

	fmt.Println("Connected to MongoDB!")
	return client
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)
	return collection
}
