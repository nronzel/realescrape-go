package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/nronzel/realescrape-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Collection, error) {

	uri := "mongodb://127.0.0.1:27017"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
        return nil, fmt.Errorf("Error connecting to DB: %v", err)
	}

	// Ping Mongo
	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, fmt.Errorf("Error pinging database: %v\n", err)
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
		return nil, fmt.Errorf("Error: %v", err)
	}

	return collection, nil
}

func InsertMongo(collection *mongo.Collection) error {
	// Read data from master.json
	data, err := ioutil.ReadFile("master.json")

	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	var houses []models.House

	err = json.Unmarshal(data, &houses)
	if err != nil {
		return fmt.Errorf("Error umarshaling json: %v\n", err)
	}

	// Filters by unique key, and updates the entry if it exists, else creates
	// a new entry.
	for _, h := range houses {
		filter := bson.M{"link": h.Link}
		update := bson.M{"$set": h}

		opts := options.Update().SetUpsert(true)

		_, err = collection.UpdateOne(context.Background(), filter, update, opts)
		if err != nil {
			return fmt.Errorf("Error updating/inserting data into collection: %v", err)
		}
	}
	return nil
}
