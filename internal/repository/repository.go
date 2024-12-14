package repository

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// UserRepository represents a repository for users and user_data tables.
type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}

// UserDataCache represents a cache for user data.
type UserDataCache interface {
	SaveField(ctx context.Context, userID int64, field string, value interface{}) error
	GetFromCache(ctx context.Context, userID int64) (*model.UserData, error)
}

// UserDataRepository represents a repository for user_data table and cache.
type UserDataRepository interface {
	UserDataCache

	Create(ctx context.Context, userData *model.UserData) (int64, error)
	Get(ctx context.Context, id int64) (*model.UserData, error)
	Update(ctx context.Context, userID int64, userData *model.UserData) (int64, error)
}
