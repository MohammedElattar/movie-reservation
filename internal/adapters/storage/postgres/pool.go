// Package postgres
package postgres

import (
	"context"
	"fmt"

	"github.com/MohammedElattar/movie-reservation/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SslMode,
	)

	cfgPool, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	return pgxpool.NewWithConfig(context.Background(), cfgPool)
}
