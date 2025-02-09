package service

import (
	events "github.com/crosleyzack/price_tracker/internal/events/fsjson"
	items "github.com/crosleyzack/price_tracker/internal/items/fsjson"
	sites "github.com/crosleyzack/price_tracker/internal/sites/fsjson"
)

type Config struct {
	Event *events.Config
	Item  *items.Config
	Site  *sites.Config
}

func NewConfig() (*Config, error) {
	conf := Config{}
	var err error
	if conf.Event, err = events.NewConfig(); err != nil {
		return nil, err
	}
	if conf.Item, err = items.NewConfig(); err != nil {
		return nil, err
	}
	if conf.Site, err = sites.NewConfig(); err != nil {
		return nil, err
	}
	return &conf, nil
}
