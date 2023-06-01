package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo() *mongo.Collection {

	uri := "mongodb://127.0.0.1/27017"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	// Ping Mongo
	err = client.Ping(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB")

	collection := client.Database("realtor").Collection("houses")

    // Create unique index to prevent duplicates
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"Link": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		panic(err)
	}

	// Create separate function to insert data

	return collection
}

func insertMongo(collection *mongo.Collection, house house) {

}
