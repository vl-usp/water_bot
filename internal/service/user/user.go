package user

import (
	"context"

	"github.com/vl-usp/water_bot/internal/constants"
	"github.com/vl-usp/water_bot/internal/model"
)

// CreateUser creates a new user.
func (s *serv) CreateUser(ctx context.Context, user model.User) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.userStore.CreateUser(ctx, user)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userStore.GetUser(ctx, user.ID)
		return errTx
	})

	return err
}

// GetUser returns a user by id.
func (s *serv) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	return s.userStore.GetUser(ctx, userID)
}

// GetFullUser returns a user by id with all params.
func (s *serv) GetFullUser(ctx context.Context, userID int64) (*model.User, error) {
	return s.userStore.GetFullUser(ctx, userID)
}

// SaveUserParam saves UserParam to cache.
func (s *serv) SaveUserParam(ctx context.Context, userID int64, field string, value interface{}) error {
	return s.userCache.SaveUserParam(ctx, userID, field, value)
}

// UpdateUserFromCache gets UserParams from cache, gets User from db and update user
func (s *serv) UpdateUserFromCache(ctx context.Context, userID int64) error {
	return s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		user, errTx := s.userStore.GetUser(ctx, userID)
		if errTx != nil {
			return errTx
		}

		cacheParams, err := s.userCache.GetUserParams(ctx, userID)
		if err != nil {
			return err
		}

		if user.Params != nil {
			dbParams, errTx := s.userStore.GetUserParams(ctx, user.Params.ID)
			if errTx != nil {
				return errTx
			}
			params := s.mergeUserParams(*dbParams, *cacheParams)
			errTx = s.userStore.UpdateUserParams(ctx, user.Params.ID, params)
			if errTx != nil {
				return errTx
			}
			return nil
		}

		fullParams, errTx := s.userStore.FillUserParams(ctx, *cacheParams)
		if errTx != nil {
			return errTx
		}
		fullParams.WaterGoal = s.calcWaterGoal(*fullParams)
		paramsID, errTx := s.userStore.CreateUserParams(ctx, *fullParams)
		if errTx != nil {
			return errTx
		}

		user.Params = &model.UserParams{
			ID: paramsID,
		}
		return s.userStore.UpdateUser(ctx, userID, *user)
	})
}

func (s *serv) calcWaterGoal(params model.UserParams) uint16 {
	if params.WaterGoal != 0 {
		return params.WaterGoal
	}

	if params.Weight == nil {
		return uint16(constants.WaterGoalDefault)
	}

	goal := constants.WaterGoalDefault + float64(*params.Weight)*params.Sex.WaterCoef*params.PhysicalActivity.WaterCoef*params.Climate.WaterCoef
	return uint16(goal)
}

func (s *serv) mergeUserParams(params1 model.UserParams, params2 model.UserParams) model.UserParams {
	if params2.Sex != nil {
		params1.Sex = params2.Sex
	}

	if params2.PhysicalActivity != nil {
		params1.PhysicalActivity = params2.PhysicalActivity
	}

	if params2.Climate != nil {
		params1.Climate = params2.Climate
	}

	if params2.Timezone != nil {
		params1.Timezone = params2.Timezone
	}

	if params2.Weight != nil {
		params1.Weight = params2.Weight
	}

	if params2.WaterGoal != 0 {
		params1.WaterGoal = s.calcWaterGoal(params2)
	}

	return params1
}
