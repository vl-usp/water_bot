package user

import (
	"github.com/vl-usp/water_bot/internal/storage"
	"github.com/vl-usp/water_bot/pkg/client/db"
)

type store struct {
	db db.Client
}

// New returns a new user repository.
func New(db db.Client) storage.User {
	return &store{
		db: db,
	}
}
