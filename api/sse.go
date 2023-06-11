package api

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func sseHandler(c echo.Context, e *echo.Echo, eb *EventBus) error {
	e.Logger.Info("SSE Handler called")
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	eventChannel := eb.Subscribe()

	for {
		select {
		case currEvent, ok := <-eventChannel:
			if !ok {
				// If the channel is closed, break the loop
				break
			}

			_, err := fmt.Fprintf(c.Response(), "data: %s\n\n", currEvent)
			if err != nil {
				e.Logger.Error(fmt.Sprintf("Error writing to response: %v", err))
				// This likely means that the client has disconnected.
				break
			}

			c.Response().Flush()
		case <-time.After(500 * time.Second):
			// Close the connection if no message is received in above time
			return nil
		}
	}
}
