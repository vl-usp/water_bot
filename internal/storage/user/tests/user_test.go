package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/storage"
	"github.com/vl-usp/water_bot/internal/storage/mocks"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx  context.Context
		user model.User
	}

	type output struct {
		err error
	}

	ctx := context.Background()
	mc := minimock.NewController(t)

	user := *model.FakeUser()

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{
			name: "success case",
			in: input{
				ctx:  ctx,
				user: user,
			},
			out: output{
				err: nil,
			},
		},
		{
			name: "storage error case",
			in: input{
				ctx:  ctx,
				user: model.User{},
			},
			out: output{
				err: fmt.Errorf("storage error"),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userStorageMock := func(mc *minimock.Controller) storage.User {
				storageMock := mocks.NewUserMock(mc)
				storageMock.CreateUserMock.Expect(ctx, tt.in.user).Return(tt.out.err)
				return storageMock
			}
			storageMock := storage.User(userStorageMock(mc))
			err := storageMock.CreateUser(tt.in.ctx, tt.in.user)
			require.Equal(t, tt.out.err, err)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx    context.Context
		userID int64
		user   model.User
	}

	type output struct {
		err error
	}

	ctx := context.Background()
	mc := minimock.NewController(t)

	user := *model.FakeUser()

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{
			name: "success case",
			in: input{
				ctx:    ctx,
				userID: user.ID,
				user:   user,
			},
			out: output{
				err: nil,
			},
		},
		{
			name: "storage error case",
			in: input{
				ctx:    ctx,
				userID: 0,
				user:   model.User{},
			},
			out: output{
				err: fmt.Errorf("storage error"),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userStorageMock := func(mc *minimock.Controller) storage.User {
				storageMock := mocks.NewUserMock(mc)
				storageMock.UpdateUserMock.Expect(ctx, tt.in.userID, tt.in.user).Return(tt.out.err)
				return storageMock
			}

			storageMock := storage.User(userStorageMock(mc))
			err := storageMock.UpdateUser(tt.in.ctx, tt.in.userID, tt.in.user)
			require.Equal(t, tt.out.err, err)
		})
	}
}

func TestGetUser(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx    context.Context
		userID int64
	}

	type output struct {
		user *model.User
		err  error
	}

	ctx := context.Background()
	mc := minimock.NewController(t)

	user := model.FakeUser()

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{
			name: "success case",
			in: input{
				ctx:    ctx,
				userID: user.ID,
			},
			out: output{
				user: user,
				err:  nil,
			},
		},
		{
			name: "storage error case",
			in: input{
				ctx:    ctx,
				userID: 0,
			},
			out: output{
				user: nil,
				err:  fmt.Errorf("storage error"),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userStorageMock := func(mc *minimock.Controller) storage.User {
				storageMock := mocks.NewUserMock(mc)
				storageMock.GetUserMock.Expect(ctx, tt.in.userID).Return(tt.out.user, tt.out.err)
				return storageMock
			}

			storageMock := storage.User(userStorageMock(mc))
			user, err := storageMock.GetUser(tt.in.ctx, tt.in.userID)
			require.Equal(t, tt.out.user, user)
			require.Equal(t, tt.out.err, err)
		})
	}
}

func TestGetFullUser(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx    context.Context
		userID int64
	}

	type output struct {
		user *model.User
		err  error
	}

	ctx := context.Background()
	mc := minimock.NewController(t)

	user := model.FakeUser()
	params := model.FakeUserParams()
	params.Sex = model.FakeSex()
	params.PhysicalActivity = model.FakePhysicalActivity()
	params.Climate = model.FakeClimate()
	params.Timezone = model.FakeTimezone()
	user.Params = params

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{
			name: "success case",
			in: input{
				ctx:    ctx,
				userID: user.ID,
			},
			out: output{
				user: user,
				err:  nil,
			},
		},
		{
			name: "storage error case",
			in: input{
				ctx:    ctx,
				userID: 0,
			},
			out: output{
				user: nil,
				err:  fmt.Errorf("storage error"),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userStorageMock := func(mc *minimock.Controller) storage.User {
				storageMock := mocks.NewUserMock(mc)
				storageMock.GetFullUserMock.Expect(ctx, tt.in.userID).Return(tt.out.user, tt.out.err)
				return storageMock
			}

			storageMock := storage.User(userStorageMock(mc))
			user, err := storageMock.GetFullUser(tt.in.ctx, tt.in.userID)
			require.Equal(t, tt.out.user, user)
			require.Equal(t, tt.out.err, err)
		})
	}
}
