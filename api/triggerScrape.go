package api

import (
	"log"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/nronzel/realescrape-go/scraper"
	"go.mongodb.org/mongo-driver/mongo"
)

func triggerScrape(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		locationParam := c.Param("location")
		log.Printf("Location parameter (encoded): %s", locationParam)
		location, err := url.QueryUnescape(locationParam)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to decode location param.")
		}
		log.Printf("Location (decoded): %s", location)

		go func() {
			err := scraper.RunScraper(collection, location)
			if err != nil {
				log.Printf("Failed to scrape data: %v", err)
			}
		}()

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Extraction started for location: " + location,
		})
	}
}
