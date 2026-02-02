package config

type DBDriver string

const (
	Postgres DBDriver = "psql"
)

type DBConfig struct {
	Default DBDriver `env:"DB_DEFAULT,required" envDefault:"psql"`
	Psql    struct {
		Host     string `env:"DB_HOST,required" envDefault:"127.0.0.1"`
		Port     int    `env:"DB_PORT,required" envDefault:"5432"`
		Username string `env:"DB_USERNAME" envDefault:"golang"`
		Password string `env:"DB_PASSWORD" envDefault:"golang"`
		SslMode  string `env:"DB_SSL_MODE" envDefault:"prefer"`
	}
}
