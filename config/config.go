package config

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type Config struct {
	ApplicationPort           string        `env:"APPLICATION_PORT"`
	DatabaseUrl               string        `env:"DB_URL"`
	DebugMode                 bool          `env:"DEBUG" envDefault:"false"`
	DatabaseHost              string        `env:"DATABASE_HOST"`
	DatabasePassword          string        `env:"DATABASE_PASSWORD"`
	DatabasePort              int           `env:"DB_PORT"`
	DatabaseConnectionTimeout time.Duration `env:"POSTGRES_DB_CONNECTION_TIMEOUT" envDefault:"15s"`

	PerPageItemLimit string `env:"PER_PAGE_ITEM_LIMIT" envDefault:"10"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}
