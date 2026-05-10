package koanfconfig

import (
	"fmt"

	"github.com/IgorSteps/dblinter/internal/domain"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// KoanfConfig holds configurations for all rules with Koanf tags.
type KoanfConfig struct {
	MaxOpenConns *MaxOpenConnsConfig `koanf:"max_open_conns"`
}

// LoadConfig loads default and user configs (if provided) and merges them.
func LoadConfig(userPath string) (KoanfConfig, error) {
	k := koanf.New(".")

	if err := k.Load(file.Provider("default.yaml"), yaml.Parser()); err != nil {
		return KoanfConfig{}, err
	}

	if userPath != "" {
		if err := k.Load(file.Provider(userPath), yaml.Parser()); err != nil {
			return KoanfConfig{}, err
		}
	}

	var cfg KoanfConfig
	if err := k.Unmarshal("", &cfg); err != nil {
		return KoanfConfig{}, err
	}

	return cfg, nil
}

// MaxOpenConnsConfig holds configurations for MaxOpenConns rule with Koanf tags.
type MaxOpenConnsConfig struct {
	Enabled  *bool `koanf:"enabled"`
	Required *int  `koanf:"required"`
}

// ToDomain converts Koanf config into domain representation.
func (s *MaxOpenConnsConfig) ToDomain() (domain.MaxOpenConnsConfig, error) {
	if s.Enabled == nil {
		return domain.MaxOpenConnsConfig{}, fmt.Errorf("max_open_conns: enabled cannot be nil")
	}

	if s.Required == nil {
		return domain.MaxOpenConnsConfig{}, fmt.Errorf("max_open_conns: required cannot be nil")
	}

	cfg, err := domain.NewMaxOpenConnsConfig(*s.Enabled, *s.Required)
	if err != nil {
		return domain.MaxOpenConnsConfig{}, fmt.Errorf("max_open_conns: %v", err)
	}

	return cfg, nil
}
