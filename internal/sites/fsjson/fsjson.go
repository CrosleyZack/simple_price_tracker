package fsjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/crosleyzack/price_tracker/internal/model"
)

// SiteStore is a store for websites
type SiteStore struct {
	Path  string
	Sites map[string]model.Website
}

// New creates a new Json ItemStore
func New(config *Config) (*SiteStore, error) {
	store := &SiteStore{
		Path:  config.FileName,
		Sites: make(map[string]model.Website),
	}
	content, err := os.ReadFile(config.FileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return store, nil
		}
		return nil, errors.Join(model.ErrFileOpen, err)
	}
	sites := make([]model.Website, 0)
	err = json.Unmarshal(content, &sites)
	if err != nil {
		return nil, errors.Join(model.ErrUnmarshal, err)
	}
	for _, site := range sites {
		store.Sites[site.Name] = site
	}
	return store, nil
}

// GetSite returns a single site from the store
func (e *SiteStore) GetSite(name string) (*model.Website, error) {
	site, ok := e.Sites[name]
	if !ok {
		return nil, model.ErrNotFound
	}
	return &site, nil
}

// ListSites returns all sites from the store
func (e *SiteStore) ListSites() ([]model.Website, error) {
	sites := make([]model.Website, 0, len(e.Sites))
	for _, site := range e.Sites {
		sites = append(sites, site)
	}
	return sites, nil
}

// AddSite adds a site to the store
func (e *SiteStore) AddSite(site model.Website) error {
	e.Sites[site.Name] = site
	err := e.save()
	if err != nil {
		return fmt.Errorf("Error during save(): %w", err)
	}
	return nil
}

func (e *SiteStore) save() error {
	sites := make([]model.Website, 0, len(e.Sites))
	for _, site := range e.Sites {
		sites = append(sites, site)
	}
	content, err := json.Marshal(sites)
	if err != nil {
		return errors.Join(model.ErrMarshal, err)
	}
	err = os.WriteFile(e.Path, content, 0644)
	if err != nil {
		return errors.Join(model.ErrFileWrite, err)
	}
	return nil
}
