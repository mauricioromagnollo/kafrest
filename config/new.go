package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// NewConfig initializes a new Config instance by processing environment variables.
// It uses the envconfig library to map environment variables to the Config struct.
// Returns a pointer to the Config instance and an error if the processing fails.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToLoadEnvironments, err)
	}

	return cfg, nil
}
