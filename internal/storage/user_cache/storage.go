package user_cache

import (
	"github.com/vl-usp/water_bot/internal/storage"
	"github.com/vl-usp/water_bot/pkg/client/cache"
)

type store struct {
	cache cache.Client
}

// New returns a new user repository.
func New(cache cache.Client) storage.UserCache {
	return &store{
		cache: cache,
	}
}
