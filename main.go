package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nronzel/realescrape-go/api"
	"github.com/nronzel/realescrape-go/db"
	"github.com/nronzel/realescrape-go/scraper"
)

func main() {
    // Connect to collection
	collection := db.ConnectMongo()

    if len(os.Args) < 2 {
        fmt.Println("Please provide a location as an argument.")
        os.Exit(1)
    }

	location := strings.ReplaceAll(os.Args[1], " ", "_")

	scraper.RunScraper(collection, location)

    // Start the Echo API on localhost:3000
    api.StartAPI(collection)
}
