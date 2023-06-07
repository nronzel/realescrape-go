package api

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nronzel/realescrape-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getAllHouses(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		var page, limit int
		var err error

		if pageParam := c.QueryParam("page"); pageParam != "" {
			page, err = strconv.Atoi(pageParam)
			if err != nil || page < 1 {
				page = 1
			}
		}

		if limitParam := c.QueryParam("limit"); limitParam != "" {
			limit, err = strconv.Atoi(limitParam)
			if err != nil || limit < 1 {
				limit = 30
			}
		}

		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		// Call the cancel function to avoid context leak
		defer cancel()

		findOptions := options.Find()
		if limit != 0 {
			findOptions.SetLimit(int64(limit))
		}
		findOptions.SetLimit(int64(limit))
		findOptions.SetSkip(int64((page - 1) * limit))

		// Attempts to find all documents in MongoDB collection
		cursor, err := collection.Find(ctx, bson.M{}, findOptions)
		if err != nil {
			// Log the error
			log.Printf("Failed to find houses: %v", err)
			// Return generic error message to client with 500 status
			return c.JSON(http.StatusInternalServerError, "Failed to find houses")
		}

		// Close the cursor if no errors found when done
		defer cursor.Close(ctx)

		houses := []models.House{}

		// Iterate over cursor, decoding each doc into House struct
		for cursor.Next(ctx) {
			var house models.House
			err := cursor.Decode(&house)
			if err != nil {
				// Log the error
				log.Printf("Failed to decode house: %v", err)
				// Return a generic error message to client with 500 status
				return c.JSON(http.StatusInternalServerError, "Failed to decode house")
			}

			// Append the house to our slice if no errors
			houses = append(houses, house)
		}

		// Check for remaining cursor errors after loop
		if err := cursor.Err(); err != nil {
			log.Printf("Cursor error: %v", err)
			return c.JSON(http.StatusInternalServerError, "Failed to retrieve houses")
		}

		// Send 200 status code with the response
		err = c.JSON(http.StatusOK, houses)
		if err != nil {
			log.Printf("Failed to send response: %v", err)
		}

		return err
	}
}
