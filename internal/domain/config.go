package domain

type Config struct {
	MaxOpenConns string
}

func NewConfig(maxOpenConns string) *Config {
	return &Config{
		MaxOpenConns: maxOpenConns,
	}
}
