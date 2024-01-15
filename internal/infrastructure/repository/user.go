package repository

import (
	"context"
	"database/sql"

	"github.com/alanpryoga/go-clean-architecture/internal/domain"
	"github.com/alanpryoga/go-clean-architecture/internal/infrastructure/cache"
)

type userRepository struct {
	db    *sql.DB
	cache cache.Cache
}

func NewUserRepository(db *sql.DB, cache cache.Cache) domain.UserRepository {
	return &userRepository{
		db:    db,
		cache: cache,
	}
}

func (*userRepository) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	panic("unimplemented")
}

func (*userRepository) Insert(ctx context.Context, user domain.User) error {
	panic("unimplemented")
}
