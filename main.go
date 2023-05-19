/*
TODO
[]lotsize conversion
[]ratios
[]git repo
more custom parameters
[]beds
[]baths
[]sqft
[]single/multi family
*/

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

type House struct {
	Price     string
	Beds      string
	Baths     string
	Sqft      string
	lotSize   string
	lotUnit   string
	Street    string
	City      string
	State     string
	Zip       string
	Link      string
	CrawlTime string
}

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

func writeHousesToCSV(houses []House) error {
	location := strings.ReplaceAll(os.Args[1], " ", "-")
	filePath := filepath.Join("scans", fmt.Sprintf("%s.csv", location))

	// Create the scans directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	headers := []string{
		"Price", "Beds", "Baths", "Sqft", "LotSize",
		"LotUnit", "Street", "City", "State", "Zip", "Link", "CrawlTime",
	}
	if err := writer.Write(headers); err != nil {
		return err
	}

    var writeErr error
	// Write data
	for _, h := range houses {
		record := []string{
			h.Price, h.Beds, h.Baths, h.Sqft, h.lotSize,
			h.lotUnit, h.Street, h.City, h.State, h.Zip, h.Link, h.CrawlTime,
		}
		if err := writer.Write(record); err != nil {
			writeErr = err
		}
	}

	return writeErr
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

func parseHouse(e *colly.HTMLElement) House {
	temp := House{}
	temp.Price = strings.Replace(e.ChildText("span[data-label='pc-price']"), "$", "", 1)
	temp.Beds = strings.TrimSuffix(e.ChildText("li[data-label='pc-meta-beds'] span"), "bed")
	temp.Baths = strings.TrimSuffix(e.ChildText("li[data-label='pc-meta-baths'] span"), "bath")
	temp.Baths = strings.ReplaceAll(temp.Baths, "+", "")
	temp.Sqft = strings.TrimSuffix(e.ChildText("li[data-label='pc-meta-sqft'] span"), "sqft")

    // Split and the size and unit for the lot
	lotSize := e.ChildText("li[data-label='pc-meta-sqftlot'] span")
	size, lotUnit := splitUnits(lotSize)
	temp.lotSize = size
	temp.lotUnit = lotUnit

    // Split the address info to provide slightly better normalization
	address := e.ChildText("div[data-label='pc-address']")
	street, city, state, zip := parseAddress(address)
	temp.Street = street
	temp.City = city
	temp.State = state
	temp.Zip = zip
	temp.Link = "https://realtor.com" + e.ChildAttr("div.photo-wrap a", "href")

	currTime := time.Now()
	dateTime := currTime.Format("2006-01-02 15:04:05")
	temp.CrawlTime = dateTime

	return temp
}

func logStats(start time.Time, houses []House) {
	log.Println("Finished Scraping page.")
	elapsed := time.Since(start)
	log.Printf("Scraped: %d listings", len(houses))
	log.Printf("Elapsed Time: %s\n", elapsed)
	if len(houses) > 0 {
		averageTimePerListing := elapsed.Seconds() / float64(len(houses))
		log.Printf("Average Time Per Listing (seconds): %f", averageTimePerListing)
	}
}

func main() {
	start := time.Now()

	location := strings.ReplaceAll(os.Args[1], " ", "_")
	options := fmt.Sprintf("%s/beds-1/baths-1/%s/%s/age-3+/pnd-hide/55p-hide/sby-6/", location, homeType, minPrice)

	url := baseURL + options

	houses := []House{}
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
