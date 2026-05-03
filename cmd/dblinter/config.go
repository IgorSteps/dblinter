package main

import (
	"github.com/IgorSteps/dblinter/internal/rules"
	"github.com/spf13/viper"
)

type Config struct {
	MaxOpenConnectionsRuleConfig rules.MaxOpenConnsRuleConfig `mapstructure:"max_open_connections_rule"`
}

func NewConfig() (Config, error) {
	var config Config
	viper.SetConfigName("bench-config")
	viper.AddConfigPath("./cmd/dblinter")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
