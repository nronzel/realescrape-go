package api

import (
	"log"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartAPI(collection *mongo.Collection) {
	// Create new Echo app
	e := echo.New()

	// Enable CORS
	e.Use(middleware.CORS())

	// Register getAllHouses handler to /houses endpoint
	// Responds to requests with all items in the MongoDB collection
	e.GET("/houses", getAllHouses(collection))

	// Start server on port localhost:3000
	err := e.Start(":3000")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
