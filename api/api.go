package api

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartAPI(collection *mongo.Collection) {
	// Create new Echo app
	e := echo.New()

	// Enable CORS
	e.Use(middleware.CORS())

	// Enable recovery
	e.Use(middleware.Recover())

	// Enable logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Assign handlers to endpoints
	e.GET("/houses", getAllHouses(collection))
	e.GET("/houses/count", getHousesCount(collection))
	e.POST("/scrape/:location", triggerScrape(collection))

	// This is extremely destructive! There is no authentication or security
	// measures in place. Hosting this publicly will allow anyone to
	// nuke the database.
	e.POST("/cleardb", cleanHouse(collection))

	// Start server on port localhost:3000
	err := e.Start(":3000")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
