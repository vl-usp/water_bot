package repository

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// UserRepository represents a user repository.
type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
