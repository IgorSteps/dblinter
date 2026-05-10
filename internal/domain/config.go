package domain

import (
	"fmt"
)

// RuntimeConfig is the runtime configuration used by db linter.
type RuntimeConfig struct {
	MaxOpenConns MaxOpenConnsConfig
}

// MaxOpenConnsConfig is the runtime configuration used by MaxOpenConns rule.
type MaxOpenConnsConfig struct {
	Enabled  bool
	Required int
}

// NewMaxOpenConnsConfig returns new runtime configuration for the MaxOpenConns rule.
func NewMaxOpenConnsConfig(enabled bool, required int) (MaxOpenConnsConfig, error) {
	if required <= 0 {
		return MaxOpenConnsConfig{}, fmt.Errorf("required must be > 0")
	}

	cfg := MaxOpenConnsConfig{
		Enabled:  enabled,
		Required: required,
	}

	return cfg, nil
}
