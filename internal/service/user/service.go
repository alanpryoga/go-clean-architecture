package user

import (
	"context"

	"github.com/alanpryoga/go-clean-architecture/internal/domain"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=user

type Service interface {
	RegisterCustomer(ctx context.Context, req RegisterCustomerRequest) (User, error)
	RegisterMerchant(ctx context.Context, req RegisterMerchantRequest) (User, error)
}

type userService struct {
	userRepo   domain.UserRepository
	walletRepo domain.WalletRepository
}

func NewService(userRepo domain.UserRepository, walletRepo domain.WalletRepository) Service {
	return &userService{
		userRepo:   userRepo,
		walletRepo: walletRepo,
	}
}

func (*userService) RegisterCustomer(ctx context.Context, req RegisterCustomerRequest) (User, error) {
	panic("unimplemented")
}

func (*userService) RegisterMerchant(ctx context.Context, req RegisterMerchantRequest) (User, error) {
	panic("unimplemented")
}
