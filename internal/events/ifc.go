package events

import (
	"github.com/crosleyzack/price_tracker/internal/model"
)

// IFC is the interface for the event store
type IFC interface {
	ListEvents() ([]model.Event, error)
	ListEventsForItem(name string) ([]model.Event, error)
	AddEvent(name string, event model.Event) error
	CurrentPrice(name string) (float32, error)
}
