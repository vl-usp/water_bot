package user

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/vl-usp/water_bot/internal/model"
	refsConverter "github.com/vl-usp/water_bot/internal/storage/refs/converter"
	modelRefRepo "github.com/vl-usp/water_bot/internal/storage/refs/model"
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
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.CreateUserParams",
		QueryRaw: query,
	}

	var id int64
	err = s.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update updates a user data by user id.
func (s *store) UpdateUserParams(ctx context.Context, id int64, params model.UserParams) error {
	builder := sq.Update(userParamsTable).
		PlaceholderFormat(sq.Dollar).
		Set(sexIDColumn, params.Sex.ID).
		Set(physicalActivityIDColumn, params.PhysicalActivity.ID).
		Set(climateIDColumn, params.Climate.ID).
		Set(timezoneIDColumn, params.Timezone.ID).
		Set(weightColumn, params.Weight).
		Set(waterGoalColumn, params.WaterGoal).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.UpdateUserParams",
		QueryRaw: query,
	}

	_, err = s.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetFullUserParams fill input params with data from ref tables.
func (s *store) GetFullUserParams(ctx context.Context, params model.UserParams) (*model.UserParams, error) {
	query := `
		select 	s.id, s.key, s.name, s.water_coef, 
				pa.id, pa.key, pa.name, pa.water_coef
			   	c.id, c.key, c.name, c.water_coef,
				tz.id, tz.name, tz.cities, tz.utc_offset
		from unnest(array[true])
		left join ref_sex s on s.id = $1
		left join ref_physical_activity pa on pa.id = $2
		left join ref_climate c on c.id = $3
		left join ref_timezone tz on tz.id = $4;
	`

	q := db.Query{
		Name:     "user_repository.GetUser",
		QueryRaw: query,
	}

	var sex modelRefRepo.Sex
	var physicalActivity modelRefRepo.PhysicalActivity
	var climate modelRefRepo.Climate
	var timezone modelRefRepo.Timezone

	err := s.db.DB().QueryRowContext(ctx, q, []byte{params.Sex.ID, params.PhysicalActivity.ID, params.Climate.ID, params.Timezone.ID}).
		Scan(
			&sex.ID, &sex.Key, &sex.Name, &sex.WaterCoef,
			&physicalActivity.ID, &physicalActivity.Key, &physicalActivity.Name, &physicalActivity.WaterCoef,
			&climate.ID, &climate.Key, &climate.Name, &climate.WaterCoef,
			&timezone.ID, &timezone.Name, &timezone.Cities, &timezone.UTCOffset,
		)
	if err != nil {
		return nil, err
	}

	params.Sex = refsConverter.ToSexFromRepo(sex)
	params.PhysicalActivity = refsConverter.ToPhysicalActivityFromRepo(physicalActivity)
	params.Climate = refsConverter.ToClimateFromRepo(climate)
	params.Timezone = refsConverter.ToTimezoneFromRepo(timezone)

	return &params, nil
}
