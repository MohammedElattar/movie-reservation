// Package config
package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
	Logger Logger

	DB struct {
		Host string `env:"DB_HOST,required" envDefault:"localhost"`
	}
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
