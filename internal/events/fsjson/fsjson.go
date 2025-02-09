package fsjson

import (
	"encoding/json"
	"errors"
	"os"
	"slices"
	"time"

	"github.com/crosleyzack/price_tracker/internal/model"
)

// EventStore is a store for events
type EventStore struct {
	Path   string
	Events map[string][]model.Event
}

// New creates a new Json ItemStore
func New(config *Config) (*EventStore, error) {
	store := &EventStore{
		Path: config.FileName,
	}
	content, err := os.ReadFile(config.FileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			store.Events = make(map[string][]model.Event)
			return store, nil
		}
		return nil, errors.Join(model.ErrFileOpen, err)
	}
	err = json.Unmarshal(content, &store.Events)
	if err != nil {
		return nil, errors.Join(model.ErrUnmarshal, err)
	}
	return store, nil
}

// ListEvents returns all stored events for all items, sorted by date
func (e *EventStore) ListEvents() ([]model.Event, error) {
	events := make([]model.Event, 0)
	for item, itemEvents := range e.Events {
		for _, event := range itemEvents {
			event.Item = item
			events = append(events, event)
		}
	}
	slices.SortFunc(events, func(i, j model.Event) int {
		return i.Date.Compare(j.Date)
	})
	return events, nil
}

// ListEventsForItem returns all stored events for a single item, sorted by date
func (e *EventStore) ListEventsForItem(name string) ([]model.Event, error) {
	events, ok := e.Events[name]
	if !ok {
		return nil, model.ErrNotFound
	}
	return events, nil
}

// AddEvent adds a new price event to the store
func (e *EventStore) AddEvent(name string, event model.Event) error {
	event.Date = time.Now().UTC()
	events, ok := e.Events[name]
	if !ok {
		events = make([]model.Event, 0)
	}
	// prepend so events are sorted by date
	events = append([]model.Event{event}, events...)
	e.Events[name] = events
	return e.save()
}

// CurrentPrice returns the current price for an item
func (e *EventStore) CurrentPrice(name string) (float32, error) {
	events, ok := e.Events[name]
	if !ok {
		return -1, model.ErrNotFound
	}
	if len(events) == 0 {
		return -1, model.ErrEmpty
	}
	return events[0].Price, nil
}

func (e *EventStore) save() error {
	content, err := json.Marshal(e.Events)
	if err != nil {
		return errors.Join(model.ErrMarshal, err)
	}
	err = os.WriteFile(e.Path, content, 0644)
	if err != nil {
		return errors.Join(model.ErrFileWrite, err)
	}
	return nil
}
