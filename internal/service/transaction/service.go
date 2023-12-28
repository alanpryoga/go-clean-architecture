package transaction

import "github.com/alanpryoga/go-clean-architecture/internal/domain"

//go:generate mockgen -source=service.go -destination=service_mock.go -package=transaction

type Service interface {
}

type transactionService struct {
	userRepo   domain.UserRepository
	walletRepo domain.WalletRepository
	trxRepo    domain.TransactionRepository
}

func NewService(
	userRepo domain.UserRepository,
	walletRepo domain.WalletRepository,
	trxRepo domain.TransactionRepository,
) Service {
	return &transactionService{
		userRepo:   userRepo,
		walletRepo: walletRepo,
		trxRepo:    trxRepo,
	}
}
