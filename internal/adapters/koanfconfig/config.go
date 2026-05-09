package koanfconfig

import (
	"fmt"

	"github.com/IgorSteps/dblinter/internal/domain"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type KoanfConfig struct {
	MaxOpenConns *MaxOpenConnsConfig `koanf:"max_open_conns"`
}

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

type MaxOpenConnsConfig struct {
	Enabled  *bool `koanf:"enabled"`
	Required *int  `koanf:"required"`
}

func (s *MaxOpenConnsConfig) ToDomain() (domain.MaxOpenConnsConfig, error) {
	cfg, err := domain.NewMaxOpenConnsConfig(*s.Enabled, *s.Required)
	if err != nil {
		return domain.MaxOpenConnsConfig{}, fmt.Errorf("max_open_conns: %v", err)
	}

	return cfg, nil
}
