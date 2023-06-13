package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/nronzel/realescrape-go/scraper"
	"go.mongodb.org/mongo-driver/mongo"
)

func triggerScrape(collection *mongo.Collection, eb *EventBus) echo.HandlerFunc {
	return func(c echo.Context) error {
		locationParam := c.Param("location")
		location, err := url.QueryUnescape(locationParam)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to decode location param.")
		}

		go func() {
			err := scraper.RunScraper(collection, location)
			if err != nil {
				c.Logger().Error(err)
				log.Printf("Failed to scrape data: %v", err)
			}
			fmt.Println("trigger scrape event sent")
			eb.Publish("db_updated")
		}()

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Extraction started for location: " + location,
		})
	}
}
