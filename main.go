package main

import (
	"os"
	"strings"

	"github.com/nronzel/realescrape-go/db"
	"github.com/nronzel/realescrape-go/scraper"
)

func main() {
    // Connect to collection
	collection := db.ConnectMongo()

	location := strings.ReplaceAll(os.Args[1], " ", "_")

	scraper.RunScraper(collection, location)
}
