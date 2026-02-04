package postgres

import (
	"github.com/MohammedElattar/movie-reservation/internal/storage/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
	pool *pgxpool.Pool
	q    *sqlc.Queries
}

func NewPostgresStore(pool *pgxpool.Pool) *PostgresStore {
	return &PostgresStore{
		pool: pool,
		q:    sqlc.New(pool),
	}
}

func (store *PostgresStore) Close() {
	store.pool.Close()
}
