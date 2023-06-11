package api

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func doCleanHouse(ctx context.Context, collection *mongo.Collection) error {
	// Delete everything in MongoDB collection
	_, err := collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		return err
	}

	// Remove master.json
	err = os.Remove("master.json")
	if err != nil {
		return err
	}

    // Read the directory
	files, err := ioutil.ReadDir("data/")
	if err != nil {
		return err
	}

    // Remove all JSON files in /data
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".json") {
			err := os.Remove("data/" + f.Name())
			if err != nil {
				return err
			}
		}
	}

	log.Println("Clean house operation completed successfully.")

	return nil
}

// Deletes all items in the MongoDB collection, as well as all json files
//
//	in the /data directory and master.json. This sets a completely clean
//	slate for the database and API.
func cleanHouse(collection *mongo.Collection, eb *EventBus) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := doCleanHouse(ctx, collection)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		eb.Publish("db_updated")

		return c.JSON(http.StatusOK, echo.Map{
			"message": "Clean house operaton completed successfully",
		})
	}
}
