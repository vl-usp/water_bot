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

func TestCreateUserParams(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx    context.Context
		params model.UserParams
	}

	type output struct {
		paramsID int64
		err      error
	}
	ctx := context.Background()
	mc := minimock.NewController(t)

	params := *model.FakeUserParams()
	params.Sex = model.FakeSex()
	params.PhysicalActivity = model.FakePhysicalActivity()
	params.Climate = model.FakeClimate()
	params.Timezone = model.FakeTimezone()

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{
			name: "success case",
			in: input{
				ctx:    ctx,
				params: params,
			},
			out: output{
				paramsID: params.ID,
				err:      nil,
			},
		},
		{
			name: "storage error case",
			in: input{
				ctx:    ctx,
				params: model.UserParams{},
			},
			out: output{
				paramsID: 0,
				err:      fmt.Errorf("storage error"),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userStorageMock := func(mc *minimock.Controller) storage.User {
				storageMock := mocks.NewUserMock(mc)
				storageMock.CreateUserParamsMock.Expect(ctx, tt.in.params).Return(tt.out.paramsID, tt.out.err)
				return storageMock
			}
			storageMock := storage.User(userStorageMock(mc))
			id, err := storageMock.CreateUserParams(tt.in.ctx, tt.in.params)
			require.Equal(t, tt.out.err, err)
			require.Equal(t, tt.out.paramsID, id)
		})
	}
}

func TestUpdateUserParams(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx      context.Context
		paramsID int64
		params   model.UserParams
	}

	type output struct {
		err error
	}
	ctx := context.Background()
	mc := minimock.NewController(t)

	params := *model.FakeUserParams()
	params.Sex = model.FakeSex()
	params.PhysicalActivity = model.FakePhysicalActivity()
	params.Climate = model.FakeClimate()
	params.Timezone = model.FakeTimezone()

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{
			name: "success case",
			in: input{
				ctx:      ctx,
				paramsID: params.ID,
				params:   params,
			},
			out: output{
				err: nil,
			},
		},
		{
			name: "storage error case",
			in: input{
				ctx:      ctx,
				paramsID: 0,
				params:   model.UserParams{},
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
				storageMock.UpdateUserParamsMock.Expect(ctx, tt.in.paramsID, tt.in.params).Return(tt.out.err)
				return storageMock
			}
			storageMock := storage.User(userStorageMock(mc))
			err := storageMock.UpdateUserParams(tt.in.ctx, tt.in.paramsID, tt.in.params)
			require.Equal(t, tt.out.err, err)
		})
	}
}

func TestGetUserParams(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx      context.Context
		paramsID int64
	}

	type output struct {
		params *model.UserParams
		err    error
	}
	ctx := context.Background()
	mc := minimock.NewController(t)

	params := model.FakeUserParams()
	params.Sex = model.FakeSex()
	params.PhysicalActivity = model.FakePhysicalActivity()
	params.Climate = model.FakeClimate()
	params.Timezone = model.FakeTimezone()

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{
			name: "success case",
			in: input{
				ctx:      ctx,
				paramsID: params.ID,
			},
			out: output{
				params: params,
				err:    nil,
			},
		},
		{
			name: "storage error case",
			in: input{
				ctx:      ctx,
				paramsID: 0,
			},
			out: output{
				params: nil,
				err:    fmt.Errorf("storage error"),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userStorageMock := func(mc *minimock.Controller) storage.User {
				storageMock := mocks.NewUserMock(mc)
				storageMock.GetUserParamsMock.Expect(ctx, tt.in.paramsID).Return(tt.out.params, tt.out.err)
				return storageMock
			}
			storageMock := storage.User(userStorageMock(mc))
			params, err := storageMock.GetUserParams(tt.in.ctx, tt.in.paramsID)
			require.Equal(t, tt.out.err, err)
			require.Equal(t, tt.out.params, params)
		})
	}
}

func TestFillUserParams(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx    context.Context
		params model.UserParams
	}

	type output struct {
		params *model.UserParams
		err    error
	}
	ctx := context.Background()
	mc := minimock.NewController(t)

	params := *model.FakeUserParams()

	filledParams := new(model.UserParams)
	*filledParams = params
	filledParams.Sex = model.FakeSex()
	filledParams.PhysicalActivity = model.FakePhysicalActivity()
	filledParams.Climate = model.FakeClimate()
	filledParams.Timezone = model.FakeTimezone()

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{
			name: "success case",
			in: input{
				ctx:    ctx,
				params: params,
			},
			out: output{
				params: filledParams,
				err:    nil,
			},
		},
		{
			name: "storage error case",
			in: input{
				ctx:    ctx,
				params: model.UserParams{},
			},
			out: output{
				params: nil,
				err:    fmt.Errorf("storage error"),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userStorageMock := func(mc *minimock.Controller) storage.User {
				storageMock := mocks.NewUserMock(mc)
				storageMock.FillUserParamsMock.Expect(ctx, tt.in.params).Return(tt.out.params, tt.out.err)
				return storageMock
			}
			storageMock := storage.User(userStorageMock(mc))
			res, err := storageMock.FillUserParams(tt.in.ctx, tt.in.params)
			require.Equal(t, tt.out.err, err)
			require.Equal(t, tt.out.params, res)
		})
	}
}
