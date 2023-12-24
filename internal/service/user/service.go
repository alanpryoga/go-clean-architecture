package user

import "github.com/alanpryoga/go-clean-architecture/internal/domain"

//go:generate mockgen -source=service.go -destination=service_mock.go -package=user

type Service interface {
}

type userService struct {
	cacheRepo domain.CacheRepository
	dbRepo    domain.DatabaseRepository
}

func NewService(cacheRepo domain.CacheRepository, dbRepo domain.DatabaseRepository) Service {
	return &userService{
		cacheRepo: cacheRepo,
		dbRepo:    dbRepo,
	}
}
