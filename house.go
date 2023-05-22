package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type house struct {
	Price     string `json:"Price"`
	Beds      string `json:"Beds"`
	Baths     string `json:"Baths"`
	Sqft      string `json:"Sqft"`
	LotSize   string `json:"LotSize"`
	LotUnit   string `json:"LotUnit"`
	LotSqft   string `json:"LotSqft"`
	Hty       string `json:"Hty"`
	HtyPcnt   string `json:"HtyPcnt"`
	Street    string `json:"Street"`
	City      string `json:"City"`
	State     string `json:"State"`
	Zip       string `json:"Zip"`
	Link      string `json:"Link"`
	CrawlTime string `json:"CrawlTime"`
}

func parseHouse(e *colly.HTMLElement) house {
	temp := house{}
	temp.Price = strings.Replace(e.ChildText("span[data-label='pc-price']"), "$", "", 1)
	temp.Beds = strings.TrimSuffix(e.ChildText("li[data-label='pc-meta-beds'] span"), "bed")
	temp.Baths = strings.TrimSuffix(e.ChildText("li[data-label='pc-meta-baths'] span"), "bath")
	temp.Baths = strings.ReplaceAll(temp.Baths, "+", "")
	temp.Sqft = strings.TrimSuffix(e.ChildText("li[data-label='pc-meta-sqft'] span"), "sqft")

	/*
	 Splits lotsize and lotsize unit, also calculates total
	 lotsize in sqft
	*/
	lotSize := e.ChildText("li[data-label='pc-meta-sqftlot'] span")
	size, lotUnit, totalSqft := splitUnits(lotSize)
	temp.LotSize = size
	temp.LotUnit = lotUnit
	temp.LotSqft = totalSqft

	// Ratios
	hty, htyPercent := htyRatios(strings.ReplaceAll(temp.Sqft, ",", ""),
		strings.ReplaceAll(temp.LotSqft, ",", ""))
	temp.Hty = hty
	temp.HtyPcnt = htyPercent

	// Split the address info to provide slightly better normalization
	address := e.ChildText("div[data-label='pc-address']")
	street, city, state, zip := parseAddress(address)
	temp.Street = street
	temp.City = city
	temp.State = state
	temp.Zip = zip
	temp.Link = "https://realtor.com" + e.ChildAttr("div.photo-wrap a", "href")

	// Add in current time & date when listing was scraped
	currTime := time.Now()
	dateTime := currTime.Format("2006-01-02 15:04:05")
	temp.CrawlTime = dateTime

	return temp
}

func writeHousesToCSV(houses []house) error {
	location := strings.ReplaceAll(os.Args[1], " ", "-")
	filePath := filepath.Join("scans", fmt.Sprintf("%s.csv", location))

	// Create the "scans" directory if it doesn't exist
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

	// Set the headers of the CSV
	headers := []string{
		"Price", "Beds", "Baths", "Sqft", "LotSize",
		"LotUnit", "LotSqft", "Hty", "HtyPcnt", "Street", "City", "State", "Zip", "Link", "CrawlTime",
	}

	// Write the headers
	if err := writer.Write(headers); err != nil {
		return err
	}

	var writeErr error

	// Write data
	for _, h := range houses {
		record := []string{
			h.Price, h.Beds, h.Baths, h.Sqft, h.LotSize,
			h.LotUnit, h.LotSqft, h.Hty, h.HtyPcnt, h.Street, h.City,
			h.State, h.Zip, h.Link, h.CrawlTime,
		}
		if err := writer.Write(record); err != nil {
			writeErr = err
		}
	}

	return writeErr
}

func writeHousesToJson(houses []house) error {
	location := strings.ReplaceAll(os.Args[1], " ", "-")
	filePath := filepath.Join("scans", fmt.Sprintf("%s.json", location))
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	jsonEncoder := json.NewEncoder(file)
	jsonEncoder.SetIndent("", "  ")
	if err := jsonEncoder.Encode(houses); err != nil {
		return err
	}

	return nil
}

func writeBothFiles(houses []house) error {

	if err := writeHousesToJson(houses); err != nil {
		return fmt.Errorf("error writing houses to JSON: %w", err)
	}

	if err := writeHousesToCSV(houses); err != nil {
		return fmt.Errorf("error writing houses to CSV: %w", err)
	}

	return nil
}
