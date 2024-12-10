package service

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// UserService represents a user service.
type UserService interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
