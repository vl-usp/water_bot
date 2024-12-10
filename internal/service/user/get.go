package user

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// Get returns a user by id.
func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
