package sites

import (
	"github.com/crosleyzack/price_tracker/internal/model"
)

// IFC is the interface for the website store
type IFC interface {
	GetSite(name string) (*model.Website, error)
	ListSites() ([]model.Website, error)
	AddSite(site model.Website) error
}
