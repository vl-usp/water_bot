package service

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// UserService represents a user service.
type UserService interface {
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	GetUser(ctx context.Context, id int64) (*model.User, error)

	UserDataService
}

// UserDataService represents a user data service.
type UserDataService interface {
	SaveUserDataField(ctx context.Context, userID int64, field string, value interface{}) error
	CreateUserData(ctx context.Context, userID int64) (*model.UserData, error)
	GetUserData(ctx context.Context, id int64) (*model.UserData, error)
	UpdateUserData(ctx context.Context, userID int64, userData *model.UserData) (int64, error)
}
