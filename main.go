package main

import (
	"fmt"

	"github.com/nronzel/realescrape-go/api"
	"github.com/nronzel/realescrape-go/db"
)

func main() {
	// Connect to collection
	collection, err := db.ConnectMongo()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Start the Echo API on localhost:3000
	api.StartAPI(collection)
}
