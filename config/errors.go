package config

import "errors"

var (
	// ErrFailedToLoadEnvironments is returned when the application fails to load the environment variables.
	ErrFailedToLoadEnvironments = errors.New("failed to load environments")
)
