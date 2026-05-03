package domain

type Config struct {
	MaxOpenConns int
}

func NewConfig(maxOpenConns int) *Config {
	return &Config{
		MaxOpenConns: maxOpenConns,
	}
}
