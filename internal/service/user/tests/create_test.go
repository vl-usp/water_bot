package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"github.com/vl-usp/water_bot/internal/client/db"
	dbMocks "github.com/vl-usp/water_bot/internal/client/db/mocks"
	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/repository"
	repoMocks "github.com/vl-usp/water_bot/internal/repository/mocks"
	"github.com/vl-usp/water_bot/internal/service/user"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req *model.User
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id         = gofakeit.Int64()
		first_name = gofakeit.FirstName()
		last_name  = gofakeit.LastName()
		username   = gofakeit.Username()
		lang_code  = "ru"
		created_at = gofakeit.Date()

		repoErr = fmt.Errorf("repo error")

		req = &model.User{
			ID:           id,
			FirstName:    first_name,
			LastName:     last_name,
			Username:     username,
			LanguageCode: lang_code,
			CreatedAt:    created_at,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		txManagerMock      txManagerMockFunc
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(t)
				return mock.ReadCommittedMock.Set(func(ctx context.Context, fn db.Handler) error {
					return fn(ctx)
				})
			},
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(id, nil)
				mock.GetMock.Expect(ctx, id).Return(req, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(t)
				return mock.ReadCommittedMock.Set(func(ctx context.Context, fn db.Handler) error {
					return fn(ctx)
				})
			},
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(0, repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			txManagerMock := tt.txManagerMock(mc)
			userRepoMock := tt.userRepositoryMock(mc)
			service := user.NewMockService(userRepoMock, txManagerMock)

			newID, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
