package fsjson

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/crosleyzack/price_tracker/internal/model"
)

// ItemStore is a store for items
type ItemStore struct {
	Path  string
	Items map[string]model.Item
}

// New creates a new Json ItemStore
func New(config *Config) (*ItemStore, error) {
	store := &ItemStore{
		Path:  config.FileName,
		Items: make(map[string]model.Item),
	}

	content, err := os.ReadFile(config.FileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return store, nil
		}
		return nil, errors.Join(model.ErrFileOpen, err)
	}

	// Now let's unmarshall the data into `payload`
	var payload []model.Item
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return nil, errors.Join(model.ErrUnmarshal, err)
	}
	for _, item := range payload {
		store.Items[item.Name] = item
	}
	return store, nil
}

// ListItems returns all stored items
func (j *ItemStore) ListItems() ([]model.Item, error) {
	items := make([]model.Item, 0, len(j.Items))
	for _, item := range j.Items {
		items = append(items, item)
	}
	return items, nil
}

// GetItem returns a single item from the store
func (j *ItemStore) GetItem(name string) (*model.Item, error) {
	item, ok := j.Items[name]
	if !ok {
		return nil, model.ErrNotFound
	}
	return &item, nil
}

// AddItem adds a new item to the store
func (j *ItemStore) AddItem(item model.Item) error {
	j.Items[item.Name] = item
	return j.save()
}

func (j *ItemStore) save() error {
	// create a list of items
	var payload []model.Item
	for _, item := range j.Items {
		payload = append(payload, item)
	}
	// create json byte array
	content, err := json.Marshal(payload)
	if err != nil {
		return errors.Join(model.ErrMarshal, err)
	}
	err = os.WriteFile(j.Path, content, 0644)
	if err != nil {
		return errors.Join(model.ErrFileWrite, err)
	}
	return nil
}
