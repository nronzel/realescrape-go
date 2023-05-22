package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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
func splitUnits(input string) (string, string, string) {
	re := regexp.MustCompile(`(\d[\d,]*(?:\.\d+)?)\s*([a-zA-Z]+)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) >= 3 {
		number := matches[1]
		unit := matches[2]
		var totalSqft string
		if unit == "acre" {
			totalSqft, err := convertToSqft(number)
			if err != nil {
				log.Println("Error converting acre to sqft.", err)
			}
			return number, unit, totalSqft
		} else {
			totalSqft = number
			return number, unit, totalSqft
		}
	}

	return "", "", ""
}

/*
Splits the address into its individual parts
I opted to keep street # and street name together
*/
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
	log.Printf("Extracted: %d listings", len(houses))
	log.Printf("Elapsed Time: %s\n", elapsed)

	if len(houses) > 0 {
		averageTimePerListing := elapsed.Seconds() / float64(len(houses))
		log.Printf("Average Time Per Listing (seconds): %f", averageTimePerListing)
	}
}

func convertToSqft(acre string) (string, error) {
	num, err := strconv.ParseFloat(acre, 32)

	if err != nil {
		return "0", fmt.Errorf("Error converting total sqft: %v", err)
	}

	if num < 0 {
		return "0", fmt.Errorf("value cannot be negative")
	}

	result := num * 43560.0

	return strconv.Itoa(int(result)), nil
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
*/
func htyRatios(houseSqft, lotSqft string) (string, string) {
	if houseSqft == "" || lotSqft == "" {
		return "", ""
	}
	houseSqftInt, err := strconv.Atoi(houseSqft)
	if err != nil {
		log.Println(err)
	}

	lotSqftInt, err := strconv.Atoi(lotSqft)
	if err != nil {
		log.Println(err)
	}

	houseSqftFloat := float64(houseSqftInt)
	lotSqftFloat := float64(lotSqftInt)

	hty := lotSqftFloat / houseSqftFloat
	htyPercent := houseSqftFloat / lotSqftFloat

	return strconv.FormatFloat(hty, 'f', 2, 64),
		strconv.FormatFloat(htyPercent, 'f', 2, 64)
}

// Combines all JSON files in the /scans dir into a single JSON
func combineJSON() {
    dir := "./scans"

    // holds json data & file count
    var data []house
    var jsonFiles int

    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatalf("failed to list files in directory %q: %v", dir, err)
    }

    // Loop through files in the directory
    for _, file := range files {
        // If file is ".json"
        if filepath.Ext(file.Name()) == ".json" {
            jsonFiles++
            // Read the contents of the file
            contents, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
            if err != nil {
                log.Printf("failed to read file %q: %v", file.Name(), err)
                continue
            }

            // Unmarshal the data
            var jsonData []house
            if err := json.Unmarshal(contents, &jsonData); err != nil {
                log.Printf("failed to unmarshal JSON data from file %q: %v", file.Name(), err)
                continue
            }


            // Append the data from this file to the master slice
            data = append(data, jsonData...)
        }
    }

    if jsonFiles <= 1 {
        os.Exit(0)
    }


    // Marshal master slice to JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Fatalf("failed to marshal data to JSON: %v", err)
    }

    // Write JSON to new file
    if err := ioutil.WriteFile("merged.json", jsonData, 0644); err != nil {
        log.Fatalf("failed to write JSON data to file: %v", err)
    }
}



