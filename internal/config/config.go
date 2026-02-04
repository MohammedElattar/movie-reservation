// Package config
package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	App         AppConfig
	Logger      Logger
	DB          DBConfig
	MaxPostSize int64
}

func Load() (*Config, error) {
	appEnv := os.Getenv("APP_ENV")

	switch appEnv {
	case "test":
		_ = godotenv.Load(".env.testing")
	case "production":
		_ = godotenv.Load(".env.production")
	default:
		_ = godotenv.Load(".env")
	}

	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, err
	}

	setConfigDefaults(&cfg)

	return &cfg, nil
}

func setConfigDefaults(cfg *Config) {
	cfg.MaxPostSize = 10 << 20 // 1mb
}
