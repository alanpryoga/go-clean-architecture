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

	cacheRepo := domain.NewMockCacheRepository(ctrl)
	dbRepo := domain.NewMockDatabaseRepository(ctrl)

	type args struct {
		cacheRepo domain.CacheRepository
		dbRepo    domain.DatabaseRepository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "Create a new Service instance",
			args: args{
				cacheRepo: cacheRepo,
				dbRepo:    dbRepo,
			},
			want: &userService{
				cacheRepo: cacheRepo,
				dbRepo:    dbRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewService(tt.args.cacheRepo, tt.args.dbRepo)

			assert.Equal(t, tt.want, got)
		})
	}
}
