package service

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// User represents a user service.
type User interface {
	CreateUser(ctx context.Context, user model.User) error
	UpdateUserFromCache(ctx context.Context, userID int64) error
	GetUser(ctx context.Context, userID int64) (*model.User, error)
	GetFullUser(ctx context.Context, userID int64) (*model.User, error)
	SaveUserParam(ctx context.Context, userID int64, field string, value interface{}) error
}

// Reference represents a reference service.
type Reference interface {
	SexList(ctx context.Context) ([]model.Sex, error)
	PhysicalActivityList(ctx context.Context) ([]model.PhysicalActivity, error)
	ClimateList(ctx context.Context) ([]model.Climate, error)
	TimezoneList(ctx context.Context) ([]model.Timezone, error)
}
