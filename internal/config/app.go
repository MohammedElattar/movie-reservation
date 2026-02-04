package config

type (
	Locale string
	AppEnv string
)

const (
	LocaleEn Locale = "en"
	LocaleAr Locale = "ar"
)

const (
	EnvDevelopment = "development"
	EnvTesting     = "testing"
	EnvProduction  = "production"
)

type AppConfig struct {
	AppName  string `env:"APP_NAME" envDefault:"GoLang"`
	AppEnv   AppEnv `env:"APP_ENV" envDefault:"development"`
	AppPort  int    `env:"APP_PORT,required" envDefault:"8080"`
	Locale   Locale `env:"APP_LOCALE" envDefault:"en"`
	FallbackLocale   Locale `env:"APP_FALLBACK_LOCALE" envDefault:"en"`
	Timezone string `env:"APP_TIMEZONE" envDefault:"UTC"`
}
