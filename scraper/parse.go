package scraper

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/nronzel/realescrape-go/models"
	"github.com/nronzel/realescrape-go/utils"
)

func parseHouse(e *colly.HTMLElement) models.House {
	temp := models.House{}
	tPrice := strings.Replace(e.ChildText("span[data-label='pc-price']"), "$", "", 1)
	tPrice = strings.ReplaceAll(tPrice, ",", "")
	temp.Price = utils.SafeAtoi(tPrice)

	tBeds := strings.TrimSuffix(e.ChildText("li[data-label='pc-meta-beds'] span"), "bed")
	temp.Beds = utils.SafeAtoi(tBeds)

	tBaths := strings.TrimSuffix(e.ChildText("li[data-label='pc-meta-baths'] span"), "bath")
	tBaths = strings.ReplaceAll(tBaths, "+", "")
	temp.Baths = utils.SafeParseFloat(tBaths, 1)

	tSqft := strings.TrimSuffix(e.ChildText("li[data-label='pc-meta-sqft'] span"), "sqft")
	tSqft = strings.ReplaceAll(tSqft, ",", "")
	temp.Sqft = utils.SafeAtoi(tSqft)

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
	hty, htyPercent := htyRatios(temp.Sqft, temp.LotSqft)
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
	temp.CrawlTime = currTime

	return temp
}

func writeHousesToCSV(houses []models.House, location string) error {
	location = strings.ReplaceAll(location, " ", "-")

	filePath := filepath.Join("data", fmt.Sprintf("%s.csv", location))

	// Create the "data" directory if it doesn't exist
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
		"LotUnit", "LotSqft", "Hty", "HtyPcnt", "Street", "City", "State",
		"Zip", "Link", "CrawlTime",
	}

	// Write the headers
	if err := writer.Write(headers); err != nil {
		return err
	}

	var writeErr error

	// Write data
	for _, h := range houses {
		crawlTimeStr := h.CrawlTime.Format("2006-01-02 15:04:05")
		priceStr := strconv.Itoa(h.Price)
		bedStr := strconv.Itoa(h.Beds)
		bathStr := fmt.Sprintf("%.1f", h.Baths)
		sqftStr := strconv.Itoa(h.Sqft)
		sizeStr := fmt.Sprintf("%.2f", h.LotSize)
		lotSqftStr := strconv.Itoa(h.LotSqft)
		htyStr := fmt.Sprintf("%.2f", h.Hty)
		htyPcntStr := fmt.Sprintf("%.2f", h.HtyPcnt)

		record := []string{
			priceStr, bedStr, bathStr, sqftStr, sizeStr,
			h.LotUnit, lotSqftStr, htyStr, htyPcntStr, h.Street, h.City,
			h.State, h.Zip, h.Link, crawlTimeStr,
		}
		if err := writer.Write(record); err != nil {
			writeErr = err
		}
	}

	return writeErr
}

func writeHousesToJson(houses []models.House, location string) error {
	filePath := filepath.Join("data", fmt.Sprintf("%s.json", location))
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

func writeBothFiles(houses []models.House, location string) error {

	if err := writeHousesToJson(houses, location); err != nil {
		return fmt.Errorf("error writing houses to JSON: %w", err)
	}

	if err := writeHousesToCSV(houses, location); err != nil {
		return fmt.Errorf("error writing houses to CSV: %w", err)
	}

	return nil
}
