package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const (
	homeType = "type-single-family-home"
	minPrice = "price-100000-na"
	baseURL  = "https://www.realtor.com/realestateandhomes-search/"
	agent    = "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/113.0"
)

func main() {
	start := time.Now()

	location := strings.ReplaceAll(os.Args[1], " ", "_")
	options := fmt.Sprintf("%s/beds-1/baths-1/%s/%s/age-3+/pnd-hide/55p-hide/sby-6/", location, homeType, minPrice)

	url := baseURL + options

	houses := []house{}
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

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	c.OnHTML("a[aria-label='Go to next page']", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")

		if nextPage != "" {
			nextPage = strings.Split(nextPage, "/")[3]
			nextPageURL := url + nextPage
			err := c.Visit(nextPageURL)
			if err != nil {
				log.Println("Error visiting next page:", err)
			}
		} else {
			logStats(start, houses)

			if err := writeHousesToCSV(houses); err != nil {
				log.Println("Error saving to CSV:", err)
			}
			os.Exit(0)
		}
	})
	c.Wait()

	logStats(start, houses)
}
