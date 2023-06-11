package api

import "sync"

type EventBus struct {
	events chan string
	once   sync.Once
}

// Create event bus
func NewEventBus() *EventBus {
	return &EventBus{
		events: make(chan string),
	}
}

// Publish message to event bus
func (eb *EventBus) Publish(event string) {
	eb.events <- event
}

// Subscribe to event bus
func (eb *EventBus) Subscribe() <-chan string {
	return eb.events
}

// Close connection
func (eb *EventBus) Close() {
	eb.once.Do(func() {
		close(eb.events)
	})
}
