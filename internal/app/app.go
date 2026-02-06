// Package app
package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/MohammedElattar/movie-reservation/internal/domain/user"
	"github.com/MohammedElattar/movie-reservation/internal/infrastructure/logger"
	"github.com/MohammedElattar/movie-reservation/internal/infrastructure/storage/postgres"
	portsLogger "github.com/MohammedElattar/movie-reservation/internal/ports/logger"
	"github.com/MohammedElattar/movie-reservation/internal/ports/storage"
	httpTransport "github.com/MohammedElattar/movie-reservation/internal/transport/http"
	"github.com/MohammedElattar/movie-reservation/pkg/i18"
	"github.com/MohammedElattar/movie-reservation/pkg/i18/ar"
	"github.com/MohammedElattar/movie-reservation/pkg/i18/en"
)

type App struct {
	Location      *time.Location
	Logger        portsLogger.Logger
	Server        *http.Server
	closableStore storage.StoreCloser
	I18           *i18.Bundle
}

func New(cfg *config.Config) (*App, error) {
	// ----------------------------
	// Register i18
	// ----------------------------

	b := i18.New()
	en.Register(b)
	ar.Register(b)

	// ----------------------------
	// Timezone
	// ----------------------------
	location, err := time.LoadLocation(cfg.App.Timezone)
	if err != nil {
		return nil, fmt.Errorf("failed to load timezone: %w", err)
	}

	// ----------------------------
	// Logger (Port + Adapter)
	// ----------------------------
	logLevel, err := mapLogLevel(cfg.Logger.Level)
	if err != nil {
		return nil, err
	}

	log := buildLogger(
		cfg.Logger.Driver,
		logLevel,
		location,
		LogTimeFormat,
	)

	// ----------------------------
	// Database Pool (Adapter)
	// ----------------------------
	pool, err := postgres.NewPool(cfg.DB.Postgres)
	if err != nil {
		return nil, err
	}

	postgresStore := postgres.NewPostgresStore(pool)

	// ----------------------------
	// JSON Response Writer
	// ----------------------------

	jsonResponse := httpTransport.NewJsonResponseWriter(b)

	// ----------------------------
	// Repositories (Adapters)
	// ----------------------------
	userRepo := postgres.NewUserRepository(postgresStore)

	// ----------------------------
	// Domain Services
	// ----------------------------
	userService := user.NewLoginService(userRepo)

	// ----------------------------
	// HTTP Transport
	// ----------------------------
	mwctx := httpTransport.NewMiddlewareContext(b, log, jsonResponse, cfg)
	userHandler := httpTransport.NewUserHandler(userService, mwctx)
	router := httpTransport.NewRouter(userHandler, mwctx)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.AppPort),
		Handler: router,
	}

	return &App{
		Logger:        log,
		Location:      location,
		Server:        server,
		closableStore: pool,
		I18:           b,
	}, nil
}

// Run starts the HTTP server and handles graceful shutdown
func (a *App) Run() error {
	a.Logger.Info(fmt.Sprintf("starting server on %s", a.Server.Addr))

	go func() {
		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.Logger.Error("http server failed", portsLogger.Error(err))
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.Logger.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	defer a.closableStore.Close()

	return a.Server.Shutdown(ctx)
}

// ----------------------------
// Logger Factory
// ----------------------------
func buildLogger(
	driver config.LoggerDriver,
	level portsLogger.Level,
	location *time.Location,
	logTimeFormat string,
) portsLogger.Logger {
	switch driver {
	case config.StdoutLogger:
		return logger.NewStdoutLogger(level, location, logTimeFormat)
	default:
		panic(fmt.Sprintf("unknown log driver %s", driver))
	}
}

// ----------------------------
// Log Level Mapping
// ----------------------------
func mapLogLevel(level config.LogLevel) (portsLogger.Level, error) {
	switch level {
	case config.DebugLevel:
		return portsLogger.DebugLevel, nil
	case config.InfoLevel:
		return portsLogger.InfoLevel, nil
	case config.WarnLevel:
		return portsLogger.WarnLevel, nil
	case config.ErrorLevel:
		return portsLogger.ErrorLevel, nil
	default:
		return portsLogger.InfoLevel, fmt.Errorf("unknown log level %s", level)
	}
}
