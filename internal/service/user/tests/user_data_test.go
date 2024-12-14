package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/repository"
	repoMocks "github.com/vl-usp/water_bot/internal/repository/mocks"
	"github.com/vl-usp/water_bot/internal/service/user"
	"github.com/vl-usp/water_bot/pkg/client/cache"
	cacheMocks "github.com/vl-usp/water_bot/pkg/client/cache/mocks"
	"github.com/vl-usp/water_bot/pkg/client/db"
	dbMocks "github.com/vl-usp/water_bot/pkg/client/db/mocks"
)

func TestCreateUserData(t *testing.T) {
	t.Parallel()
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager
	type cacheMockFunc func(mc *minimock.Controller) cache.Client
	type userDataRepositoryMockFunc func(mc *minimock.Controller) repository.UserDataRepository

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id                 = gofakeit.Int64()
		userID             = gofakeit.Int64()
		sexID              = model.Male
		physicalActivityID = model.Moderate
		climateID          = model.Temperate
		weight             = gofakeit.Int()
		waterGoal          = gofakeit.Int()
		createdAt          = gofakeit.Date()

		repoErr = fmt.Errorf("repo error")

		userData = &model.UserData{
			ID:                 id,
			UserID:             userID,
			SexID:              sexID,
			PhysicalActivityID: physicalActivityID,
			ClimateID:          climateID,
			Weight:             weight,
			WaterGoal:          waterGoal,
			CreatedAt:          createdAt,
		}
	)

	tests := []struct {
		name                   string
		args                   args
		want                   *model.UserData
		err                    error
		txManagerMock          txManagerMockFunc
		cacheMock              cacheMockFunc
		userDataRepositoryMock userDataRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: userData,
			err:  nil,
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				return mock.ReadCommittedMock.Set(func(ctx context.Context, fn db.Handler) error {
					return fn(ctx)
				})
			},
			cacheMock: func(mc *minimock.Controller) cache.Client {
				mock := cacheMocks.NewClientMock(mc)
				return mock
			},
			userDataRepositoryMock: func(mc *minimock.Controller) repository.UserDataRepository {
				mock := repoMocks.NewUserDataRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, userData).Return(id, nil)
				mock.GetFromCacheMock.Expect(ctx, id).Return(userData, nil)
				mock.GetMock.Expect(ctx, id).Return(userData, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: nil,
			err:  repoErr,
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				return mock.ReadCommittedMock.Set(func(ctx context.Context, fn db.Handler) error {
					return fn(ctx)
				})
			},
			cacheMock: func(mc *minimock.Controller) cache.Client {
				mock := cacheMocks.NewClientMock(mc)
				return mock
			},
			userDataRepositoryMock: func(mc *minimock.Controller) repository.UserDataRepository {
				mock := repoMocks.NewUserDataRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, userData).Return(0, repoErr)
				mock.GetFromCacheMock.Expect(ctx, id).Return(userData, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			txManagerMock := tt.txManagerMock(mc)
			userRepoMock := tt.userDataRepositoryMock(mc)
			cacheMock := tt.cacheMock(mc)
			service := user.NewMockService(userRepoMock, cacheMock, txManagerMock)

			userData, err := service.CreateUserData(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, userData)
		})
	}
}

func TestGetUserData(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		firstName = gofakeit.FirstName()
		lastName  = gofakeit.LastName()
		username  = gofakeit.Username()
		langCode  = "ru"
		createdAt = gofakeit.Date()

		repoErr = fmt.Errorf("repo error")

		res = &model.User{
			ID:           id,
			FirstName:    firstName,
			LastName:     lastName,
			Username:     username,
			LanguageCode: langCode,
			CreatedAt:    createdAt,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               *model.User
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: res,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(res, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: nil,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := tt.userRepositoryMock(mc)
			service := user.NewMockService(userRepoMock)

			newID, err := service.GetUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
