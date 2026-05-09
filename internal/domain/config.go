package domain

import (
	"fmt"
)

type RuntimeConfig struct {
	MaxOpenConns MaxOpenConnsConfig
}

type MaxOpenConnsConfig struct {
	Enabled  bool
	Required int
}

func NewDefaultMaxOpenConnsConfig() MaxOpenConnsConfig {
	return MaxOpenConnsConfig{
		Enabled:  true,
		Required: 10,
	}
}

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
