package user

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

	type args struct {
		userRepo   domain.UserRepository
		walletRepo domain.WalletRepository
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
			},
			want: &userService{
				userRepo:   userRepo,
				walletRepo: walletRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewService(tt.args.userRepo, tt.args.walletRepo)

			assert.Equal(t, tt.want, got)
		})
	}
}
