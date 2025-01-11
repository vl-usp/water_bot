package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"github.com/vl-usp/water_bot/internal/constants"
	"github.com/vl-usp/water_bot/internal/model"
	userService "github.com/vl-usp/water_bot/internal/service/user"
	storageMocks "github.com/vl-usp/water_bot/internal/storage/mocks"
	"github.com/vl-usp/water_bot/pkg/client/db"
	dbMocks "github.com/vl-usp/water_bot/pkg/client/db/mocks"
)

// TestCreateUser tests the CreateUser function of the user service
func TestCreateUser(t *testing.T) {
	t.Parallel()
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager
	type userStorerMockFunc func(mc *minimock.Controller) userService.Storer

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
			userStorerMock: func(mc *minimock.Controller) userService.Storer {
				mock := storageMocks.NewUserMock(mc)
				mock.CreateUserMock.Expect(ctx, req).Return(nil)
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
			userStorerMock: func(mc *minimock.Controller) userService.Storer {
				mock := storageMocks.NewUserMock(mc)
				mock.CreateUserMock.Expect(ctx, req).Return(storageErr)
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
			service := userService.NewMockService(userStorerMock, txManagerMock)

			err := service.CreateUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}

// TestGetUser tests the GetUser function of the user service
func TestGetUser(t *testing.T) {
	t.Parallel()
	type userStorerMockFunc func(mc *minimock.Controller) userService.Storer

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
			userStorerMock: func(mc *minimock.Controller) userService.Storer {
				mock := storageMocks.NewUserMock(mc)
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
			userStorerMock: func(mc *minimock.Controller) userService.Storer {
				mock := storageMocks.NewUserMock(mc)
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
			service := userService.NewMockService(userStorerMock)

			newID, err := service.GetUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}

// TestGetFullUser tests the GetUser function of the user service
func TestGetFullUser(t *testing.T) {
	t.Parallel()
	type userStorerMockFunc func(mc *minimock.Controller) userService.Storer

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

		params = model.UserParams{
			ID:  gofakeit.Int64(),
			Sex: &model.Sex{ID: 1},
			PhysicalActivity: &model.PhysicalActivity{
				ID: 1,
			},
			Climate: &model.Climate{
				ID: 1,
			},
			Timezone: &model.Timezone{
				ID: 1,
			},
		}

		storageErr = fmt.Errorf("storage error")

		res = &model.User{
			ID:           id,
			FirstName:    firstName,
			LastName:     lastName,
			Username:     username,
			LanguageCode: langCode,
			Params:       &params,
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
			userStorerMock: func(mc *minimock.Controller) userService.Storer {
				mock := storageMocks.NewUserMock(mc)
				mock.GetFullUserMock.Expect(ctx, id).Return(res, nil)
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
			userStorerMock: func(mc *minimock.Controller) userService.Storer {
				mock := storageMocks.NewUserMock(mc)
				mock.GetFullUserMock.Expect(ctx, id).Return(nil, storageErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userStorerMock := tt.userStorerMock(mc)
			service := userService.NewMockService(userStorerMock)

			newID, err := service.GetFullUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}

// TestSaveUserParam tests the SaveUserParam function of the user service
func TestSaveUserParam(t *testing.T) {
	t.Parallel()
	type userCacherMockFunc func(mc *minimock.Controller) userService.Cacher

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
			userCacherMock: func(mc *minimock.Controller) userService.Cacher {
				mock := storageMocks.NewUserCacheMock(mc)
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
			userCacherMock: func(mc *minimock.Controller) userService.Cacher {
				mock := storageMocks.NewUserCacheMock(mc)
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
			service := userService.NewMockService(userCacherMock)

			err := service.SaveUserParam(tt.args.ctx, tt.args.userID, tt.args.field, tt.args.value)
			require.Equal(t, tt.err, err)
		})
	}
}

// TestUpdateUserFromCache tests the UpdateUserFromCache function of the user service
// func TestUpdateUserFromCache(t *testing.T) {
// 	t.Parallel()
// 	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager
// 	type userStorerMockFunc func(mc *minimock.Controller) userService.Storer
// 	type userCacherMockFunc func(mc *minimock.Controller) userService.Cacher

// 	type args struct {
// 		ctx    context.Context
// 		userID int64
// 	}

// 	var (
// 		ctx = context.Background()
// 		mc  = minimock.NewController(t)

// 		fakeUser             = model.FakeUser()
// 		fakeParams           = model.FakeUserParams()
// 		fakeSex              = model.FakeSex()
// 		fakePhysicalActivity = model.FakePhysicalActivity()
// 		fakeClimate          = model.FakeClimate()
// 		fakeTimezone         = model.FakeTimezone()

// 		dbParams = &model.UserParams{
// 			ID:               fakeParams.ID,
// 			Sex:              fakeSex,
// 			PhysicalActivity: fakePhysicalActivity,
// 			Climate:          fakeClimate,
// 			Timezone:         fakeTimezone,
// 			Weight:           fakeParams.Weight,
// 			WaterGoal:        fakeParams.WaterGoal,
// 			CreatedAt:        fakeParams.CreatedAt,
// 			UpdatedAt:        fakeParams.UpdatedAt,
// 		}

// 		user = &model.User{
// 			ID:           fakeUser.ID,
// 			FirstName:    fakeUser.FirstName,
// 			LastName:     fakeUser.LastName,
// 			Username:     fakeUser.Username,
// 			LanguageCode: fakeUser.LanguageCode,
// 			Params:       &model.UserParams{ID: fakeParams.ID},
// 			CreatedAt:    fakeUser.CreatedAt,
// 		}

// 		cacheParams = &model.UserParams{
// 			Sex:              &model.Sex{ID: fakeSex.ID},
// 			PhysicalActivity: &model.PhysicalActivity{ID: fakePhysicalActivity.ID},
// 			Climate:          &model.Climate{ID: fakeClimate.ID},
// 			Timezone:         &model.Timezone{ID: fakeTimezone.ID},
// 			Weight:           fakeParams.Weight,
// 		}

// 		filledParams = &model.UserParams{
// 			Sex:              fakeSex,
// 			PhysicalActivity: fakePhysicalActivity,
// 			Climate:          fakeClimate,
// 			Timezone:         fakeTimezone,
// 			Weight:           fakeParams.Weight,
// 		}

// 		fullParams = &model.UserParams{
// 			Sex:              fakeSex,
// 			PhysicalActivity: fakePhysicalActivity,
// 			Climate:          fakeClimate,
// 			Timezone:         fakeTimezone,
// 			Weight:           fakeParams.Weight,
// 			WaterGoal:        fakeParams.WaterGoal,
// 		}

// 		// storageErr = fmt.Errorf("storage error")
// 	)

// 	tests := []struct {
// 		name           string
// 		args           args
// 		err            error
// 		txManagerMock  txManagerMockFunc
// 		userStorerMock userStorerMockFunc
// 		userCacherMock userCacherMockFunc
// 	}{
// 		{
// 			name: "success case",
// 			args: args{
// 				ctx:    ctx,
// 				userID: user.ID,
// 			},
// 			err: nil,
// 			txManagerMock: func(mc *minimock.Controller) db.TxManager {
// 				mock := dbMocks.NewTxManagerMock(mc)
// 				return mock.ReadCommittedMock.Set(func(ctx context.Context, fn db.Handler) error {
// 					return fn(ctx)
// 				})
// 			},
// 			userStorerMock: func(mc *minimock.Controller) userService.Storer {
// 				mock := storageMocks.NewUserMock(mc)
// 				mock.GetUserMock.Expect(ctx, user.ID).Return(user, nil)
// 				mock.GetUserParamsMock.Expect(ctx, user.Params.ID).Return(dbParams, nil)
// 				mock.UpdateUserParamsMock.Expect(ctx, user.Params.ID, *cacheParams).Return(nil)
// 				mock.FillUserParamsMock.Expect(ctx, *cacheParams).Return(filledParams, nil)
// 				mock.CreateUserParamsMock.Expect(ctx, *fullParams).Return(dbParams.ID, nil)
// 				mock.UpdateUserMock.Expect(ctx, user.ID, *user).Return(nil)
// 				return mock
// 			},
// 			userCacherMock: func(mc *minimock.Controller) userService.Cacher {
// 				mock := storageMocks.NewUserCacheMock(mc)
// 				mock.GetUserParamsMock.Expect(ctx, user.ID).Return(cacheParams, nil)
// 				return mock
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			txManagerMock := tt.txManagerMock(mc)
// 			userStorerMock := tt.userStorerMock(mc)
// 			userCacherMock := tt.userCacherMock(mc)
// 			service := userService.NewMockService(txManagerMock, userStorerMock, userCacherMock)

// 			t.Logf("Test: %s, user: %+v", tt.name, tt.args)

// 			err := service.UpdateUserFromCache(tt.args.ctx, tt.args.userID)
// 			require.Equal(t, tt.err, err)
// 		})
// 	}
// }
