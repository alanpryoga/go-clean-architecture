package repository

import (
	"context"
	"database/sql"

	"github.com/alanpryoga/go-clean-architecture/internal/domain"
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) domain.TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (*transactionRepository) Insert(ctx context.Context, transaction domain.Transaction) error {
	panic("unimplemented")
}
