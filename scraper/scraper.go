package scraper

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/nronzel/realescrape-go/db"
	"github.com/nronzel/realescrape-go/models"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
    USER AGENTS: To be used at a later time to pick one randomly each runtime.

Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36
(KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36

Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36
(KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.42

Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36
(KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.3

Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15
(KHTML, like Gecko) Version/16.2 Safari/605.1.1

Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36
(KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36

Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0

Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36
(KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36

Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15
(KHTML, like Gecko) Version/16.4 Safari/605.1.1

*/

const (
	homeType = "type-single-family-home"
	minPrice = "price-100000-na"
	radius   = "radius-1"
	baseURL  = "https://www.realtor.com/realestateandhomes-search/"
	agent    = "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/113.0"
)

func RunScraper(collection *mongo.Collection, location string) error {
	start := time.Now()

	decodedLocation, err := url.QueryUnescape(location)
    decodedLocation = strings.ReplaceAll(decodedLocation, " ", "_")

	if err != nil {
		return fmt.Errorf("Failed to decode location: %v", err)
	}

	options := fmt.Sprintf(
		"%s/beds-1/baths-1/%s/%s/age-3+/pnd-hide/fc-hide/55p-hide/%s/sby-6/",
		decodedLocation,
		homeType,
		minPrice,
		radius,
	)

	url := baseURL + options

	houses := []models.House{}
	c := getCollector()

	c.OnHTML("li[data-testid='result-card']", func(e *colly.HTMLElement) {
		houses = append(houses, parseHouse(e))
	})

	// max parallelism & introduce random delay
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// Before making request, print
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("a[aria-label='Go to next page']", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")

		if nextPage != "" {
			nextPage = strings.Split(nextPage, "/")[3]
			nextPageURL := url + nextPage
			err := c.Visit(nextPageURL)
			if err != nil {
				log.Println("Error visiting next page:", err)
			}
		}
	})

	err = c.Visit(url)

	if err != nil {
		return fmt.Errorf("Error visiting URL: %v", err)
	}

	c.Wait()

	logStats(start, houses)

    // Writes CSV and JSON file named with location
	if err := writeBothFiles(houses, location); err != nil {
		return fmt.Errorf("Error while writing files: %v", err)
	}

	// Combine all JSON files located in /data
	if err := combineJSON(); err != nil {
		return fmt.Errorf("Error combining JSON files: %v", err)
	}

	db.InsertMongo(collection)

	return nil
}
