package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type house struct {
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

func parseHouse(e *colly.HTMLElement) house {
	temp := house{}
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

func writeHousesToCSV(houses []house) error {
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
