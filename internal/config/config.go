package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Config holds configurations for all rules.
type Config struct {
	MaxOpenConns *MaxOpenConnsConfig `koanf:"max_open_conns"`
}

// MaxOpenConnsConfig holds configurations for MaxOpenConns rule.
type MaxOpenConnsConfig struct {
	Enabled  bool `koanf:"enabled"`
	Required int  `koanf:"required"`
}

// LoadConfig loads default and user configs (if provided) and merges them and returns runtime config.
func LoadConfig(userPath string) (Config, error) {
	k := koanf.New(".")

	if err := k.Load(file.Provider("default.yaml"), yaml.Parser()); err != nil {
		return Config{}, err
	}

	if userPath != "" {
		if err := k.Load(file.Provider(userPath), yaml.Parser()); err != nil {
			return Config{}, err
		}
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
