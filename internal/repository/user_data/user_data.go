package user_data

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/repository/user_data/converter"
	modelRepo "github.com/vl-usp/water_bot/internal/repository/user_data/model"
	"github.com/vl-usp/water_bot/pkg/client/db"
)

const (
	tableName = "user_data"

	idColumn                 = "id"
	userIDColumn             = "user_id"
	sexIDColumn              = "sex_id"
	weightColumn             = "weight"
	physicalActivityIDColumn = "physical_activity_id"
	waterGoalColumn          = "water_goal"
	climateIDColumn          = "climate_id"
	createdAtColumn          = "created_at"
	updatedAtColumn          = "updated_at"
)

// Create creates a new user data.
func (r *repo) Create(ctx context.Context, userData *model.UserData) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(userIDColumn, sexIDColumn, weightColumn, physicalActivityIDColumn, climateIDColumn, waterGoalColumn).
		Values(userData.UserID, userData.SexID, userData.Weight, userData.PhysicalActivityID, userData.ClimateID, userData.WaterGoal).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_data_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get returns a user data by id.
func (r *repo) Get(ctx context.Context, id int64) (*model.UserData, error) {
	builder := sq.Select(
		idColumn,
		userIDColumn,
		sexIDColumn,
		weightColumn,
		physicalActivityIDColumn,
		climateIDColumn,
		createdAtColumn,
		updatedAtColumn,
	).PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_data_repository.Get",
		QueryRaw: query,
	}

	var userData modelRepo.UserData
	err = r.db.DB().ScanOneContext(ctx, &userData, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserDataFromRepo(&userData), nil
}

// Update updates a user data by user id.
func (r *repo) Update(ctx context.Context, userID int64, userData *model.UserData) (int64, error) {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(weightColumn, userData.Weight).
		Set(physicalActivityIDColumn, userData.PhysicalActivityID).
		Set(climateIDColumn, userData.ClimateID).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{userIDColumn: userID}).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_data_repository.Update",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
