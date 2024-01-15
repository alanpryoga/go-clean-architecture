package transaction

import (
	"context"

	"github.com/alanpryoga/go-clean-architecture/internal/domain"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=transaction

type Service interface {
	TopUp(ctx context.Context, req TopUpRequest) error
	Payment(ctx context.Context, req PaymentRequest) error
	Withdraw(ctx context.Context, req WithdrawRequest) error
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

func (*transactionService) Payment(ctx context.Context, req PaymentRequest) error {
	panic("unimplemented")
}

func (*transactionService) TopUp(ctx context.Context, req TopUpRequest) error {
	panic("unimplemented")
}

func (*transactionService) Withdraw(ctx context.Context, req WithdrawRequest) error {
	panic("unimplemented")
}
