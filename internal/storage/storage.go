package storage

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// User represents a repository for user tables.
type User interface {
	CreateUser(ctx context.Context, user model.User) error
	UpdateUser(ctx context.Context, userID int64, user model.User) error
	GetUser(ctx context.Context, userID int64) (*model.User, error)
	GetFullUser(ctx context.Context, userID int64) (*model.User, error)

	CreateUserParams(ctx context.Context, params model.UserParams) (int64, error)
	UpdateUserParams(ctx context.Context, paramsID int64, params model.UserParams) error
	GetUserParams(ctx context.Context, paramsID int64) (*model.UserParams, error)
	FillUserParams(ctx context.Context, params model.UserParams) (*model.UserParams, error)
}

// Reference represents a repository for user ref tables.
type Reference interface {
	SexList(ctx context.Context) ([]model.Sex, error)
	PhysicalActivityList(ctx context.Context) ([]model.PhysicalActivity, error)
	ClimateList(ctx context.Context) ([]model.Climate, error)
	TimezoneList(ctx context.Context) ([]model.Timezone, error)
}

// UserCache represents a cache for user params, that needed for temporary storage of user data.
type UserCache interface {
	SaveUserParam(ctx context.Context, userID int64, field string, value interface{}) error
	GetUserParams(ctx context.Context, userID int64) (*model.UserParams, error)
}
