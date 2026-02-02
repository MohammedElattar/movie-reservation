package config

type (
	LoggerDriver string
	LogLevel     string
)

const (
	StdoutLogger LoggerDriver = "stdout"
	ZapLogger    LoggerDriver = "zap"
)

const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
)

type Logger struct {
	Driver LoggerDriver `env:"LOG_DRIVER" envDefault:"stdout"`
	Level  LogLevel     `env:"LOG_LEVEL" envDefault:"info"`
}
