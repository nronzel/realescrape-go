package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Deletes all items in the MongoDB collection, as well as all json files
//
//	in the /data directory and master.json. This sets a completely clean
//	slate for the database and API.
func cleanHouse(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)
		// Delete everything in MongoDB collection
		_, err = collection.DeleteMany(c.Request().Context(), bson.D{})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Remove master.json
		err = os.Remove("master.json")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Remove all JSON files in /data
		files, err := ioutil.ReadDir("data/")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".json") {
				err := os.Remove("data/" + f.Name())
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
			}
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "Clean house operation completed successfully.",
		})
	}
}
