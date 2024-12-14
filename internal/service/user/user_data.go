package user

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// SaveUserDataField saves a user data field to cache.
func (s *serv) SaveUserDataField(ctx context.Context, userID int64, field string, value interface{}) error {
	err := s.userDataRepo.SaveField(ctx, userID, field, value)
	if err != nil {
		return err
	}

	return nil
}

// CreateUserData creates a new user data from cache.
func (s *serv) CreateUserData(ctx context.Context, userID int64) (*model.UserData, error) {
	userData, err := s.userDataRepo.GetFromCache(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Calculate water goal
	if userData.WaterGoal == 0 {
		var baseWaterNorm int
		switch userData.SexID {
		case model.Male:
			baseWaterNorm = 35
		case model.Female:
			baseWaterNorm = 30
		}

		var physicalActivityNorm int
		switch userData.PhysicalActivityID {
		case model.Low:
			physicalActivityNorm = 0
		case model.Moderate:
			physicalActivityNorm = 500
		case model.High:
			physicalActivityNorm = 1000
		}

		var climateNorm int
		switch userData.ClimateID {
		case model.Cold:
			climateNorm = 0
		case model.Temperate:
			climateNorm = 200
		case model.Warm:
			climateNorm = 400
		case model.Hot:
			climateNorm = 600
		}

		userData.WaterGoal = baseWaterNorm*userData.Weight + physicalActivityNorm + climateNorm
		// logger.Get("service", "s.CreateUserData").Debug("calculated water goal", "userData", userData)
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx := s.userDataRepo.Create(ctx, userData)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userDataRepo.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return userData, err
}

// GetUserData returns a user data by id.
func (s *serv) GetUserData(ctx context.Context, id int64) (*model.UserData, error) {
	userData, err := s.userDataRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

// UpdateUserData updates a user data by user id.
func (s *serv) UpdateUserData(ctx context.Context, userID int64, userData *model.UserData) (int64, error) {
	id, err := s.userDataRepo.Update(ctx, userID, userData)
	if err != nil {
		return 0, err
	}

	return id, nil
}
