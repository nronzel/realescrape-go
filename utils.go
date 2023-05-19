package main

import (
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func getCollector() *colly.Collector {

	c := colly.NewCollector(colly.Async(true))
	c.IgnoreRobotsTxt = true
	c.UserAgent = agent

	c.OnError(func(r *colly.Response, err error) {
		log.Println(r.Request.URL)
		log.Println("Error: ", err)
	})

	return c
}

// Splits the provided string where numbers meet letters
func splitUnits(input string) (string, string) {
	re := regexp.MustCompile(`(\d[\d,]*(?:\.\d+)?)\s*([a-zA-Z]+)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) >= 3 {
		number := matches[1]
		unit := matches[2]
		return number, unit
	}

	return "", ""
}

// Splits the address into its individual parts
// I opted to keep street # and street name together
func parseAddress(address string) (string, string, string, string) {
	splitAddress := strings.Split(address, ",")

	if len(splitAddress) < 3 {
		return "", "", "", ""
	}
	street := strings.TrimSpace(splitAddress[0])
	city := strings.TrimSpace(splitAddress[1])
	stateAndZip := strings.Split(strings.TrimSpace(splitAddress[2]), " ")
	state := strings.TrimSpace(stateAndZip[0])
	zip := strings.TrimSpace(stateAndZip[1])

	return street, city, state, zip
}

func logStats(start time.Time, houses []house) {
	log.Println("Finished Scraping page.")
	elapsed := time.Since(start)
	log.Printf("Scraped: %d listings", len(houses))
	log.Printf("Elapsed Time: %s\n", elapsed)
	if len(houses) > 0 {
		averageTimePerListing := elapsed.Seconds() / float64(len(houses))
		log.Printf("Average Time Per Listing (seconds): %f", averageTimePerListing)
	}
}

