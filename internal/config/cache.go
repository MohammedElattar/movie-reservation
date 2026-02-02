package config

type CacheStore string

const (
	Memory CacheStore = "memory"
	Redis  CacheStore = "redis"
)

type CacheConfig struct {
	DefaultStore CacheStore `env:"CACHE_STORE" envDefault:"memory"`
	Prefix       string     `env:"CACHE_PREFIX,expand"`
	Memory       struct {
		TTL int `env:"CACHE_MEMORY_TTL" envDefault:"60"`
	}
	Redis struct {
		Host     string `env:"CACHE_REDIS_HOST" envDefault:"127.0.0.1"`
		Port     int    `env:"CACHE_REDIS_PORT" envDefault:"6379"`
		Password string `env:"CACHE_REDIS_PASSWORD"`
		DB       int    `env:"CACHE_REDIS_DATABASE" envDefault:"0"`
		TTL      int    `env:"CACHE_REDIS_TTL" envDefault:"300"`
	}
}
