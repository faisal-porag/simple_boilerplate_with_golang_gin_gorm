package config

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type Config struct {
	ApplicationPort           string        `env:"APPLICATION_PORT"`
	DebugMode                 bool          `env:"DEBUG" envDefault:"false"`
	DatabaseUser              string        `env:"DB_USER"`
	DatabaseHost              string        `env:"DB_HOST"`
	DatabasePassword          string        `env:"DB_PASSWORD"`
	DatabasePort              string        `env:"DB_PORT"`
	DatabaseDbName            string        `env:"DB_NAME"`
	DatabaseConnectionTimeout time.Duration `env:"POSTGRES_DB_CONNECTION_TIMEOUT" envDefault:"15s"`
	PerPageItemLimit          string        `env:"PER_PAGE_ITEM_LIMIT" envDefault:"10"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}
