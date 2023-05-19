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

/*
    USER AGENTS: To be used at a later time to pick one randomly each runtime.

Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36
(KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246

Mozilla/5.0 (X11; CrOS x86_64 8172.45.0) AppleWebKit/537.36
(KHTML, like Gecko) Chrome/51.0.2704.64 Safari/537.36

Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9
(KHTML, like Gecko) Version/9.0.2 Safari/601.3.9

Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko)
Chrome/47.0.2526.111 Safari/537.36

Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1
*/

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
