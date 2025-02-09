package fsjson

import (
	"fmt"

	"github.com/Netflix/go-env"
)

// Config configuration for the server
type Config struct {
	FileName string `env:"SITE_FILE_NAME,default=data/sites.json"`
}

// NewConfig creates a new Config object
func NewConfig() (*Config, error) {
	// Load service config
	conf := Config{}
	if _, err := env.UnmarshalFromEnviron(&conf); err != nil {
		return nil, fmt.Errorf("Failed to initialize server configuration: %w", err)
	}
	return &conf, nil
}
