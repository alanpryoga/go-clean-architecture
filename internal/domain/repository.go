package domain

import "context"

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=domain

type CacheRepository interface {
	FindUserByPhone(ctx context.Context, phone string) (User, error)
	FindWalletByUserPhone(ctx context.Context, phone string) (Wallet, error)
	InsertWalletByUserPhone(ctx context.Context, phone string, wallet Wallet) error
	DeleteWalletByUserPhone(ctx context.Context, phone string) error
}

type DatabaseRepository interface {
	InsertUser(ctx context.Context, user User) error
	InsertWallet(ctx context.Context, wallet Wallet) error
	FindUserByPhone(ctx context.Context, phone string) (User, error)
	FindWalletByUserPhone(ctx context.Context, phone string) (Wallet, error)
	UpdateWallet(ctx context.Context, wallet Wallet) error
	InsertTransaction(ctx context.Context, transaction Transaction) error
}
