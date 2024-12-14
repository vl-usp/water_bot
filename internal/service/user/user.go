package user

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// CreateUser creates a new user.
func (s *serv) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepo.Create(ctx, user)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userRepo.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, err
}

// GetUser returns a user by id.
func (s *serv) GetUser(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
