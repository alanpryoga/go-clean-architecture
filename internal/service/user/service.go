package user

import "github.com/alanpryoga/go-clean-architecture/internal/domain"

//go:generate mockgen -source=service.go -destination=service_mock.go -package=user

type Service interface {
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
