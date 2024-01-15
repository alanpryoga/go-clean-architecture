package repository

import (
	"context"
	"database/sql"

	"github.com/alanpryoga/go-clean-architecture/internal/domain"
	"github.com/alanpryoga/go-clean-architecture/internal/infrastructure/cache"
)

type walletRepository struct {
	db    *sql.DB
	cache cache.Cache
}

func NewWalletRepository(db *sql.DB, cache cache.Cache) domain.WalletRepository {
	return &walletRepository{
		db:    db,
		cache: cache,
	}
}

func (*walletRepository) Insert(ctx context.Context, wallet domain.Wallet) error {
	panic("unimplemented")
}

func (*walletRepository) FindByUserID(ctx context.Context, id string) (domain.Wallet, error) {
	panic("unimplemented")
}

func (*walletRepository) Update(ctx context.Context, wallet domain.Wallet) error {
	panic("unimplemented")
}
