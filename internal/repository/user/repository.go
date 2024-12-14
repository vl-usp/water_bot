package user

import (
	"github.com/vl-usp/water_bot/internal/repository"
	"github.com/vl-usp/water_bot/pkg/client/db"
)

type repo struct {
	db db.Client
}

// NewRepository returns a new user repository.
func NewRepository(db db.Client) repository.UserRepository {
	return &repo{
		db: db,
	}
}
