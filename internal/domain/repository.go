package domain

import (
	"context"
)

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=domain

type UserRepository interface {
	Insert(ctx context.Context, user User) error
	FindByPhone(ctx context.Context, phone string) (User, error)
}

type WalletRepository interface {
	Insert(ctx context.Context, wallet Wallet) error
	FindByUserID(ctx context.Context, id string) (Wallet, error)
	Update(ctx context.Context, wallet Wallet) error
}

type TransactionRepository interface {
	Insert(ctx context.Context, transaction Transaction) error
}
