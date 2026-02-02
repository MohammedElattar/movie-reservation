package config

type Locale string

const (
	LocaleEn Locale = "en"
	LocaleAr Locale = "ar"
)

type AppConfig struct {
	AppName  string       `env:"APP_NAME" envDefault:"GoLang"`
	AppEnv   string       `env:"APP_ENV" envDefault:"development"`
	AppPort  int          `env:"APP_PORT,required" envDefault:"8080"`
	Locale   Locale `env:"APP_LOCALE" envDefault:"en"`
	Timezone string       `env:"APP_TIMEZONE" envDefault:"UTC"`
}
