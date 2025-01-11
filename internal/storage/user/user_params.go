package user

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/vl-usp/water_bot/internal/model"
	refsConverter "github.com/vl-usp/water_bot/internal/storage/refs/converter"
	modelRefRepo "github.com/vl-usp/water_bot/internal/storage/refs/model"
	"github.com/vl-usp/water_bot/internal/storage/user/converter"
	storageModel "github.com/vl-usp/water_bot/internal/storage/user/model"
	"github.com/vl-usp/water_bot/pkg/client/db"
)

// Create creates a new user data.
func (s *store) CreateUserParams(ctx context.Context, params model.UserParams) (int64, error) {
	builder := sq.Insert(userParamsTable).
		PlaceholderFormat(sq.Dollar).
		Columns(
			sexIDColumn,
			physicalActivityIDColumn,
			climateIDColumn,
			timezoneIDColumn,
			weightColumn,
			waterGoalColumn,
		).
		Values(
			params.Sex.ID,
			params.PhysicalActivity.ID,
			params.Climate.ID,
			params.Timezone.ID,
			params.Weight,
			params.WaterGoal,
		).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build create user params query: %w", err)
	}

	q := db.Query{
		Name:     "user_storage.CreateUserParams",
		QueryRaw: query,
	}

	var id int64
	err = s.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create user params: %w", err)
	}

	return id, nil
}

// Update updates a user data by user id.
func (s *store) UpdateUserParams(ctx context.Context, paramsID int64, params model.UserParams) error {
	builder := sq.Update(userParamsTable).
		PlaceholderFormat(sq.Dollar).
		Set(sexIDColumn, params.Sex.ID).
		Set(physicalActivityIDColumn, params.PhysicalActivity.ID).
		Set(climateIDColumn, params.Climate.ID).
		Set(timezoneIDColumn, params.Timezone.ID).
		Set(weightColumn, params.Weight).
		Set(waterGoalColumn, params.WaterGoal).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: paramsID})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build update user params query: %w", err)
	}

	q := db.Query{
		Name:     "user_storage.UpdateUserParams",
		QueryRaw: query,
	}

	_, err = s.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to update user params: %w", err)
	}

	return nil
}

// GetUserParams returns a user data by user id.
func (s *store) GetUserParams(ctx context.Context, paramsID int64) (*model.UserParams, error) {
	builder := sq.Select(
		idColumn,
		sexIDColumn,
		physicalActivityIDColumn,
		climateIDColumn,
		timezoneIDColumn,
		weightColumn,
		waterGoalColumn,
		createdAtColumn,
		updatedAtColumn,
	).
		PlaceholderFormat(sq.Dollar).
		From(userParamsTable).
		Where(sq.Eq{idColumn: paramsID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build get user params query: %w", err)
	}

	q := db.Query{
		Name:     "user_storage.GetUserParams",
		QueryRaw: query,
	}

	var params storageModel.UserParams
	err = s.db.DB().ScanOneContext(ctx, &params, q, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get user params: %w", err)
	}

	return converter.ToUserParamsFromStorage(
		params,
		modelRefRepo.Sex{ID: params.SexID},
		modelRefRepo.PhysicalActivity{ID: params.PhysicalActivityID},
		modelRefRepo.Climate{ID: params.ClimateID},
		modelRefRepo.Timezone{ID: params.TimezoneID},
	), nil
}

// FillUserParams fill input params with data from ref tables.
func (s *store) FillUserParams(ctx context.Context, params model.UserParams) (*model.UserParams, error) {
	query := `
		select 	s.id, s.key, s.name, s.water_coef, 
				pa.id, pa.key, pa.name, pa.water_coef,
			   	c.id, c.key, c.name, c.water_coef,
				tz.id, tz.name, tz.cities, tz.utc_offset
		from unnest(array[true])
		left join ref_sex s on s.id = $1
		left join ref_physical_activity pa on pa.id = $2
		left join ref_climate c on c.id = $3
		left join ref_timezone tz on tz.id = $4;
	`

	q := db.Query{
		Name:     "user_storage.FillUserParams",
		QueryRaw: query,
	}

	var sex modelRefRepo.Sex
	var physicalActivity modelRefRepo.PhysicalActivity
	var climate modelRefRepo.Climate
	var timezone modelRefRepo.Timezone

	err := s.db.DB().QueryRowContext(ctx, q, params.Sex.ID, params.PhysicalActivity.ID, params.Climate.ID, params.Timezone.ID).
		Scan(
			&sex.ID, &sex.Key, &sex.Name, &sex.WaterCoef,
			&physicalActivity.ID, &physicalActivity.Key, &physicalActivity.Name, &physicalActivity.WaterCoef,
			&climate.ID, &climate.Key, &climate.Name, &climate.WaterCoef,
			&timezone.ID, &timezone.Name, &timezone.Cities, &timezone.UTCOffset,
		)
	if err != nil {
		return nil, fmt.Errorf("failed to fill user params: %w", err)
	}

	params.Sex = refsConverter.ToSexFromStorage(sex)
	params.PhysicalActivity = refsConverter.ToPhysicalActivityFromStorage(physicalActivity)
	params.Climate = refsConverter.ToClimateFromStorage(climate)
	params.Timezone = refsConverter.ToTimezoneFromStorage(timezone)

	return &params, nil
}
