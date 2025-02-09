package items

import (
	"github.com/crosleyzack/price_tracker/internal/model"
)

// IFC is the interface for the items store
type IFC interface {
	ListItems() ([]model.Item, error)
	GetItem(name string) (*model.Item, error)
	AddItem(item model.Item) error
}
