package user

import (
	"context"

	"github.com/vl-usp/water_bot/internal/constants"
	"github.com/vl-usp/water_bot/internal/model"
)

// CreateUser creates a new user.
func (s *serv) CreateUser(ctx context.Context, user model.User) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userStore.CreateUser(ctx, user)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userStore.GetUser(ctx, id)
		return errTx
	})

	if err != nil {
		return 0, err
	}

	return id, err
}

// GetUser returns a user by id.
func (s *serv) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	return s.userStore.GetUser(ctx, userID)
}

// SaveUserParam saves UserParam to cache.
func (s *serv) SaveUserParam(ctx context.Context, userID int64, field string, value interface{}) error {
	return s.userCache.SaveUserParam(ctx, userID, field, value)
}

// UpdateUserFromCache gets UserParams from cache, gets User from db and update user
func (s *serv) UpdateUserFromCache(ctx context.Context, userID int64) error {
	params, err := s.userCache.GetUserParams(ctx, userID)
	if err != nil {
		return err
	}

	// calc waterGoal if needed
	if params.WaterGoal == 0 {
		params, err = s.userStore.GetFullUserParams(ctx, *params)
		if err != nil {
			return err
		}

		goal := constants.WaterGoalDefault + float64(params.Weight)*params.Sex.WaterCoef*params.PhysicalActivity.WaterCoef*params.Climate.WaterCoef
		params.WaterGoal = int(goal)
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		paramsID, errTx := s.userStore.CreateUserParams(ctx, *params)
		if errTx != nil {
			return errTx
		}

		user, errTx := s.userStore.GetUser(ctx, userID)
		if errTx != nil {
			return errTx
		}

		user.Params.ID = paramsID
		errTx = s.userStore.UpdateUser(ctx, userID, *user)
		return errTx
	})
	return err
}
