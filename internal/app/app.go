// Package app
package app

import (
	"fmt"
	"time"

	"github.com/MohammedElattar/movie-reservation/internal/app/logger"
	"github.com/MohammedElattar/movie-reservation/internal/config"
)

type App struct {
	Location *time.Location
	Logger   logger.Logger
}

func New(cfg *config.Config) (*App, error) {
	// Construct time location based on timezone

	location, err := time.LoadLocation(cfg.App.Timezone)
	if err != nil {
		return nil, fmt.Errorf("failed to construct time location %v", err)
	}

	// Construct logger

	logLevel, err := mapLogLevel(cfg.Logger.Level)
	if err != nil {
		return nil, err
	}

	logger, err := logger.CreateNewLogger(
		cfg.Logger.Driver,
		logLevel,
		location,
		LogTimeFormat,
	)
	if err != nil {
		return nil, err
	}

	return &App{
		Logger:   logger,
		Location: location,
	}, nil
}

func (app *App) Run() error {
	return nil
}

func mapLogLevel(level config.LogLevel) (logger.LogLevel, error) {
	switch level {
	case config.DebugLevel:
		return logger.DebugLevel, nil
	case config.InfoLevel:
		return logger.InfoLevel, nil
	case config.WarnLevel:
		return logger.WarnLevel, nil
	case config.ErrorLevel:
		return logger.ErrorLevel, nil
	default:
		return logger.InfoLevel, fmt.Errorf("unknown log level %s", level)
	}
}
