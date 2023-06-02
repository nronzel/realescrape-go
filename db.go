package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo() *mongo.Collection {

    uri := "mongodb://127.0.0.1:27017"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Println("Error connecting to DB", err)
	}

	// Ping Mongo
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Printf("Error pinging database: %v\n", err)
	}

	fmt.Println("Connected to MongoDB")

	collection := client.Database("realtor").Collection("houses")

	// Create unique index to prevent duplicates
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"link": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	return collection
}

func insertMongo(collection *mongo.Collection) {
	// Read data from master.json
	data, err := ioutil.ReadFile("master.json")

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	var houses []house

	err = json.Unmarshal(data, &houses)
	if err != nil {
		log.Printf("Error umarshaling json: %v\n", err)
	}

	housesInterface := make([]interface{}, len(houses))
	for i, v := range houses {
		housesInterface[i] = v
	}

	_, err = collection.InsertMany(context.Background(), housesInterface)
	if err != nil {
		log.Printf("Error inserting data into collection: %v", err)
	}
}
