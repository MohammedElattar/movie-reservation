package postgres

import (
	"context"

	"github.com/MohammedElattar/movie-reservation/internal/domain/user"
)

type UserRepository struct {
	store *PostgresStore
}

func NewUserRepository(ps *PostgresStore) *UserRepository {
	return &UserRepository{store: ps}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	return nil
}
