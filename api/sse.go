package api

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func sseHandler(c echo.Context, collection *mongo.Collection, e *echo.Echo) error {
	e.Logger.Info("SSE Handler called")
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	prevCount := int64(0)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		currCount, err := collection.CountDocuments(ctx, bson.M{})
		if err != nil {
			if cmdErr, ok := err.(mongo.CommandError); ok {
				e.Logger.Error(fmt.Sprintf("Error counting documents: %v. Code: %v, Name: %v", cmdErr.Message, cmdErr.Code, cmdErr.Name))
			} else {
				e.Logger.Error("Error counting documents:", err)
			}
			time.Sleep(time.Second * 2)
			continue
		}

		if prevCount != currCount {
			_, err := fmt.Fprintf(c.Response(), "data: %d\n\n", currCount)
			e.Logger.Info(fmt.Sprintf("Updated Count: %v", currCount))

			if err != nil {
				e.Logger.Error(fmt.Sprintf("Error writing to response: %v", err))
				// This likely means that the client has disconnected.
				break
			}

			c.Response().Flush()
			prevCount = currCount
		}

		time.Sleep(time.Second * 2)
	}

	return nil

}
