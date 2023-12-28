package transaction

import (
	"testing"

	"github.com/alanpryoga/go-clean-architecture/internal/domain"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := domain.NewMockUserRepository(ctrl)
	walletRepo := domain.NewMockWalletRepository(ctrl)
	trxRepo := domain.NewMockTransactionRepository(ctrl)

	type args struct {
		userRepo   domain.UserRepository
		walletRepo domain.WalletRepository
		trxRepo    domain.TransactionRepository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "Create a new Service instance",
			args: args{
				userRepo:   userRepo,
				walletRepo: walletRepo,
				trxRepo:    trxRepo,
			},
			want: &transactionService{
				userRepo:   userRepo,
				walletRepo: walletRepo,
				trxRepo:    trxRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewService(tt.args.userRepo, tt.args.walletRepo, tt.args.trxRepo)

			assert.Equal(t, tt.want, got)
		})
	}
}
