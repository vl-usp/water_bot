package user_data

import (
	"github.com/vl-usp/water_bot/internal/repository"
	"github.com/vl-usp/water_bot/pkg/client/cache"
	"github.com/vl-usp/water_bot/pkg/client/db"
)

type repo struct {
	db    db.Client
	cache cache.Client
}

// NewRepository returns a new user repository.
func NewRepository(db db.Client, cache cache.Client) repository.UserDataRepository {
	return &repo{
		db:    db,
		cache: cache,
	}
}
