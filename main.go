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

	// if len(os.Args) < 2 {
	// 	fmt.Println("Please provide a location as an argument.")
	// 	os.Exit(1)
	// }
	//
	// location := strings.ReplaceAll(os.Args[1], " ", "_")
	//
	// scraper.RunScraper(collection, location)

	// Start the Echo API on localhost:3000
	api.StartAPI(collection)
}
