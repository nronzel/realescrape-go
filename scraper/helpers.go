package scraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/nronzel/realescrape-go/models"
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
func splitUnits(input string) (float64, string, int) {
	re := regexp.MustCompile(`(\d[\d,]*(?:\.\d+)?)\s*([a-zA-Z]+)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) >= 3 {
		sqft := matches[1]
		unit := matches[2]
		if unit == "acre" {
			totalSqft, err := convertToSqft(sqft)
			if err != nil {
				log.Println("Error converting acre to sqft.", err)
			}
            sqft = strings.ReplaceAll(sqft, ",", "")
			intNumber, err := strconv.ParseFloat(sqft, 2)
			if err != nil {
                fmt.Printf("Error converting sqft to float: %v", err)
			}
			return intNumber, unit, totalSqft
		} else {
            sqft = strings.ReplaceAll(sqft, ",", "")
			intNumber, err := strconv.Atoi(sqft)
			if err != nil {
                fmt.Printf("Error converting sqft to int: %v", err)
			}
			totalSqft := intNumber
			return float64(intNumber), unit, totalSqft
		}
	}

	return 0, "", 0
}

/*
Splits the address into its individual parts
I opted to keep street # and street name together
*/
func parseAddress(address string) (string, string, string, int) {
	splitAddress := strings.Split(address, ",")

	if len(splitAddress) < 3 {
		return "", "", "", 0
	}
	street := strings.TrimSpace(splitAddress[0])
	city := strings.TrimSpace(splitAddress[1])
	stateAndZip := strings.Split(strings.TrimSpace(splitAddress[2]), " ")
	state := strings.TrimSpace(stateAndZip[0])
	zip := strings.TrimSpace(stateAndZip[1])
	intZip, err := strconv.Atoi(zip)
	if err != nil {
		fmt.Printf("Error converting zip to int: %v", err)
	}

	return street, city, state, intZip
}

func logStats(start time.Time, houses []models.House) {
	log.Println("Finished Scraping page.")
	elapsed := time.Since(start)
	log.Printf("Extracted: %d listings", len(houses))
	log.Printf("Elapsed Time: %s\n", elapsed)

	if len(houses) > 0 {
		averageTimePerListing := elapsed.Seconds() / float64(len(houses))
		log.Printf("Average Time Per Listing (seconds): %f", averageTimePerListing)
	}
}

func convertToSqft(acre string) (int, error) {
	num, err := strconv.ParseFloat(acre, 32)

	if err != nil {
		return 0, fmt.Errorf("Error converting total sqft: %v", err)
	}

	if num < 0 {
		return 0, fmt.Errorf("value cannot be negative")
	}

	result := num * 43560.0

	return int(result), nil
}

/*
Calculates house to yard ratios.

The first ratio is how many times
the house will fit in the lotsize, and the second ratio is the
percentage of the lot the house takes up.

I realize that a house may have more sqft than lotsize if it is a taller
building like an apartment or condo building. Ideally these get filtered out
when searching for single family homes only which makes these ratios actually
mean something.

There are also issues if the data in the listing itself isn't properly uploaded.
Some houses do not have sqft listed or sometimes they put the lot size as the
sqft of the home.
*/
func htyRatios(houseSqft, lotSqft int) (float64, float64) {
	if houseSqft == 0 || lotSqft == 0 {
		return 0, 0
	}

	houseSqftFloat := float64(houseSqft)
	lotSqftFloat := float64(lotSqft)

	hty := lotSqftFloat / houseSqftFloat
	htyPercent := houseSqftFloat / lotSqftFloat

    hty = math.Round(hty * 100) / 100
    htyPercent = math.Round(htyPercent * 100) / 100

	return hty, htyPercent
}

// Combines all JSON files in the /data dir into a single JSON
func combineJSON() error {
	dir := "./data"

	// holds json data & file count
	var data []models.House
	var jsonFiles int

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to list files in directory %q: %v", dir, err)
	}

	// Loop through files in the directory
	for _, file := range files {
		// If file is ".json"
		if filepath.Ext(file.Name()) == ".json" {
			jsonFiles++
			// Read the contents of the file
			contents, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return fmt.Errorf("failed to read file %q: %v", file.Name(), err)
			}

			// Unmarshal the data
			var jsonData []models.House
			if err := json.Unmarshal(contents, &jsonData); err != nil {
				return fmt.Errorf("failed to unmarshal JSON data from file %q: %v", file.Name(), err)
			}

			// Append the data from this file to the master slice
			data = append(data, jsonData...)
		}
	}

	if jsonFiles < 1 {
		return nil
	}

	// Marshal master slice to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %v", err)
	}

	// Write JSON to new file
	if err := ioutil.WriteFile("master.json", jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write JSON data to file: %v", err)
	}

	return nil
}
