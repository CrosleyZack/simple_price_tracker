package model

import (
	"fmt"
	"time"
)

// Event defines a price event for an item
type Event struct {
	Price float32   `json:"price"`
	Date  time.Time `json:"date"`
	Item  string    `json:"item"`
}

func (e Event) String() string {
	return e.Item + " : " + e.Date.Format("2006-01-02") + ": $" + fmt.Sprintf("%.2f", e.Price)
}

// EventList is a map of items to their events
type EventList map[string][]Event
