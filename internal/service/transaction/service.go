package transaction

import "github.com/alanpryoga/go-clean-architecture/internal/domain"

//go:generate mockgen -source=service.go -destination=service_mock.go -package=transaction

type Service interface {
}

type transactionService struct {
	cacheRepo domain.CacheRepository
	dbRepo    domain.DatabaseRepository
}

func NewService(cacheRepo domain.CacheRepository, dbRepo domain.DatabaseRepository) Service {
	return &transactionService{
		cacheRepo: cacheRepo,
		dbRepo:    dbRepo,
	}
}
