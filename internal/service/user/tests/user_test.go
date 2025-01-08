package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"github.com/vl-usp/water_bot/internal/constants"
	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/service/user"
	"github.com/vl-usp/water_bot/internal/storage"
	repoMocks "github.com/vl-usp/water_bot/internal/storage/mocks"
	"github.com/vl-usp/water_bot/pkg/client/db"
	dbMocks "github.com/vl-usp/water_bot/pkg/client/db/mocks"
)

// TestCreateUser tests the CreateUser function of the user service
func TestCreateUser(t *testing.T) {
	t.Parallel()
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager
	type userStorerMockFunc func(mc *minimock.Controller) storage.User

	type args struct {
		ctx context.Context
		req model.User
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

		storageErr = fmt.Errorf("storage error")

		req = model.User{
			ID:           id,
			FirstName:    firstName,
			LastName:     lastName,
			Username:     username,
			LanguageCode: langCode,
			CreatedAt:    createdAt,
		}
	)

	tests := []struct {
		name           string
		args           args
		want           int64
		err            error
		txManagerMock  txManagerMockFunc
		userStorerMock userStorerMockFunc
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
				mock := dbMocks.NewTxManagerMock(mc)
				return mock.ReadCommittedMock.Set(func(ctx context.Context, fn db.Handler) error {
					return fn(ctx)
				})
			},
			userStorerMock: func(mc *minimock.Controller) storage.User {
				mock := repoMocks.NewUserMock(mc)
				mock.CreateUserMock.Expect(ctx, req).Return(id, nil)
				mock.GetUserMock.Expect(ctx, id).Return(&req, nil)
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
			err:  storageErr,
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				return mock.ReadCommittedMock.Set(func(ctx context.Context, fn db.Handler) error {
					return fn(ctx)
				})
			},
			userStorerMock: func(mc *minimock.Controller) storage.User {
				mock := repoMocks.NewUserMock(mc)
				mock.CreateUserMock.Expect(ctx, req).Return(0, storageErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			txManagerMock := tt.txManagerMock(mc)
			userStorerMock := tt.userStorerMock(mc)
			service := user.NewMockService(userStorerMock, txManagerMock)

			newID, err := service.CreateUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}

// TestGetUser tests the GetUser function of the user service
func TestGetUser(t *testing.T) {
	t.Parallel()
	type userStorerMockFunc func(mc *minimock.Controller) storage.User

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

		storageErr = fmt.Errorf("storage error")

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
		name           string
		args           args
		want           *model.User
		err            error
		userStorerMock userStorerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: res,
			err:  nil,
			userStorerMock: func(mc *minimock.Controller) storage.User {
				mock := repoMocks.NewUserMock(mc)
				mock.GetUserMock.Expect(ctx, id).Return(res, nil)
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
			err:  storageErr,
			userStorerMock: func(mc *minimock.Controller) storage.User {
				mock := repoMocks.NewUserMock(mc)
				mock.GetUserMock.Expect(ctx, id).Return(nil, storageErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userStorerMock := tt.userStorerMock(mc)
			service := user.NewMockService(userStorerMock)

			newID, err := service.GetUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}

// TestSaveUserParam tests the SaveUserParam function of the user service
func TestSaveUserParam(t *testing.T) {
	t.Parallel()
	type userCacherMockFunc func(mc *minimock.Controller) storage.UserCache

	type args struct {
		ctx    context.Context
		userID int64
		field  string
		value  interface{}
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		keys = []string{
			constants.SexKey,
			constants.PhysicalActivityKey,
			constants.ClimateKey,
			constants.TimezoneKey,
		}

		values = []string{"value1", "value2", "value3"}

		userID = gofakeit.Int64()
		field  = gofakeit.RandomString(keys)
		value  = gofakeit.RandomString(values)

		storageErr = fmt.Errorf("storage error")
	)

	tests := []struct {
		name           string
		args           args
		err            error
		userCacherMock userCacherMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:    ctx,
				userID: userID,
				field:  field,
				value:  value,
			},
			err: nil,
			userCacherMock: func(mc *minimock.Controller) storage.UserCache {
				mock := repoMocks.NewUserCacheMock(mc)
				mock.SaveUserParamMock.Expect(ctx, userID, field, value).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx:    ctx,
				userID: userID,
				field:  field,
				value:  value,
			},
			err: storageErr,
			userCacherMock: func(mc *minimock.Controller) storage.UserCache {
				mock := repoMocks.NewUserCacheMock(mc)
				mock.SaveUserParamMock.Expect(ctx, userID, field, value).Return(storageErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userCacherMock := tt.userCacherMock(mc)
			service := user.NewMockService(userCacherMock)

			err := service.SaveUserParam(tt.args.ctx, tt.args.userID, tt.args.field, tt.args.value)
			require.Equal(t, tt.err, err)
		})
	}
}

// TestUpdateUserFromCache tests the UpdateUserFromCache function of the user service
func TestUpdateUserFromCache(t *testing.T) {
	t.Parallel()
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager
	type userStorerMockFunc func(mc *minimock.Controller) storage.User
	type userCacherMockFunc func(mc *minimock.Controller) storage.UserCache

	type args struct {
		ctx    context.Context
		userID int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		userID = gofakeit.Int64()

		paramsID = gofakeit.Int64()

		params = &model.UserParams{
			Sex: model.Sex{
				ID: 1,
			},
			PhysicalActivity: model.PhysicalActivity{
				ID: 1,
			},
			Climate: model.Climate{
				ID: 1,
			},
			Timezone: model.Timezone{
				ID: 1,
			},
			Weight: 70,
		}

		fullParams = &model.UserParams{
			Sex: model.Sex{
				ID:        1,
				Key:       "male",
				Name:      "Мужчина",
				WaterCoef: 1.0,
			},
			PhysicalActivity: model.PhysicalActivity{
				ID:        1,
				Key:       "low",
				Name:      "Низкая",
				WaterCoef: 1.0,
			},
			Climate: model.Climate{
				ID:        1,
				Key:       "cold",
				Name:      "Холодный",
				WaterCoef: 1.0,
			},
			Timezone: model.Timezone{
				ID:        1,
				Name:      "UTC+0",
				Cities:    gofakeit.City(),
				UTCOffset: 0,
			},
			Weight:    70,
			WaterGoal: 2070,
		}

		userObj = &model.User{
			ID:           userID,
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			Username:     gofakeit.Username(),
			LanguageCode: "ru",
			Params: model.UserParams{
				ID:               paramsID,
				Sex:              fullParams.Sex,
				PhysicalActivity: fullParams.PhysicalActivity,
				Climate:          fullParams.Climate,
				Timezone:         fullParams.Timezone,
				Weight:           fullParams.Weight,
				WaterGoal:        2070,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
			CreatedAt: time.Now(),
		}

		// storageErr = fmt.Errorf("storage error")
	)

	tests := []struct {
		name           string
		args           args
		err            error
		txManagerMock  txManagerMockFunc
		userStorerMock userStorerMockFunc
		userCacherMock userCacherMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			err: nil,
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				return mock.ReadCommittedMock.Set(func(ctx context.Context, fn db.Handler) error {
					return fn(ctx)
				})
			},
			userStorerMock: func(mc *minimock.Controller) storage.User {
				mock := repoMocks.NewUserMock(mc)
				mock.GetFullUserParamsMock.Expect(ctx, *params).Return(fullParams, nil)
				mock.CreateUserParamsMock.Expect(ctx, *fullParams).Return(paramsID, nil)
				mock.GetUserMock.Expect(ctx, userID).Return(userObj, nil)
				mock.UpdateUserMock.Expect(ctx, userID, *userObj).Return(nil)
				return mock
			},
			userCacherMock: func(mc *minimock.Controller) storage.UserCache {
				mock := repoMocks.NewUserCacheMock(mc)
				mock.GetUserParamsMock.Expect(ctx, userID).Return(params, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			txManagerMock := tt.txManagerMock(mc)
			userStorerMock := tt.userStorerMock(mc)
			userCacherMock := tt.userCacherMock(mc)
			service := user.NewMockService(txManagerMock, userStorerMock, userCacherMock)

			err := service.UpdateUserFromCache(tt.args.ctx, tt.args.userID)
			require.Equal(t, tt.err, err)
		})
	}
}
