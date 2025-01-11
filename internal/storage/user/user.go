package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/vl-usp/water_bot/internal/model"
	storageRefsModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
	"github.com/vl-usp/water_bot/internal/storage/user/converter"
	storageModel "github.com/vl-usp/water_bot/internal/storage/user/model"
	"github.com/vl-usp/water_bot/pkg/client/db"
	"github.com/vl-usp/water_bot/pkg/logger"
)

// Create creates a new user.
func (s *store) CreateUser(ctx context.Context, user model.User) error {
	logger.Get("storage", "CreateUser").Info("creating user...", "user", user)
	builder := sq.Insert(usersTable).
		PlaceholderFormat(sq.Dollar).
		Columns(
			idColumn,
			firstNameColumn,
			lastNameColumn,
			usernameColumn,
			languageCodeColumn,
		).
		Values(
			user.ID,
			user.FirstName,
			user.LastName,
			user.Username,
			user.LanguageCode,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build create user query: %w", err)
	}

	q := db.Query{
		Name:     "user_storage.CreateUser",
		QueryRaw: query,
	}

	_, err = s.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// UpdateUser updates a user data by user id.
func (s *store) UpdateUser(ctx context.Context, userID int64, user model.User) error {
	builder := sq.Update(usersTable).
		PlaceholderFormat(sq.Dollar).
		Set(firstNameColumn, user.FirstName).
		Set(lastNameColumn, user.LastName).
		Set(usernameColumn, user.Username).
		Set(languageCodeColumn, user.LanguageCode).
		Set(paramsIDColumn, user.Params.ID).
		Where(sq.Eq{idColumn: userID})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build update user query: %w", err)
	}

	q := db.Query{
		Name:     "user_storage.UpdateUser",
		QueryRaw: query,
	}

	_, err = s.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (s *store) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	builder := sq.Select(
		idColumn,
		firstNameColumn,
		lastNameColumn,
		usernameColumn,
		languageCodeColumn,
		paramsIDColumn,
		createdAtColumn,
	).
		PlaceholderFormat(sq.Dollar).
		From(usersTable).
		Where(sq.Eq{idColumn: userID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build get user query: %w", err)
	}

	q := db.Query{
		Name:     "user_storage.GetUser",
		QueryRaw: query,
	}

	var user storageModel.User
	err = s.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return converter.ToUserFromStorage(user), nil
}

// Get returns a user by id.
// It joins user, user_params and refs tables and return full user data.
func (s *store) GetFullUser(ctx context.Context, userID int64) (*model.User, error) {
	query := `
		select 
			u.id, u.first_name, u.last_name, u.username, u.language_code, u.created_at,
			up.id, up.weight, up.water_goal, up.created_at, up.updated_at,
			s.id, s.key, s.name, s.water_coef,
			pa.id, pa.key, pa.name, pa.water_coef,
			c.id, c.key, c.name, c.water_coef,
			tz.id, tz.name, tz.cities, tz.utc_offset
		from users u
		left join user_params up on u.params_id = up.id
		left join ref_sex s on up.sex_id = s.id
		left join ref_physical_activity pa on up.physical_activity_id = pa.id
		left join ref_climate c on up.climate_id = c.id
		left join ref_timezone tz on up.timezone_id = tz.id
		where u.id = $1	
	`

	q := db.Query{
		Name:     "user_storage.GetFullUser",
		QueryRaw: query,
	}

	var user storageModel.User
	var userParams storageModel.UserParams
	var sex storageRefsModel.Sex
	var physicalActivity storageRefsModel.PhysicalActivity
	var climate storageRefsModel.Climate
	var timezone storageRefsModel.Timezone

	err := s.db.DB().QueryRowContext(ctx, q, userID).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.LanguageCode, &user.CreatedAt,
		&userParams.ID, &userParams.Weight, &userParams.WaterGoal, &userParams.CreatedAt, &userParams.UpdatedAt,
		&sex.ID, &sex.Key, &sex.Name, &sex.WaterCoef,
		&physicalActivity.ID, &physicalActivity.Key, &physicalActivity.Name, &physicalActivity.WaterCoef,
		&climate.ID, &climate.Key, &climate.Name, &climate.WaterCoef,
		&timezone.ID, &timezone.Name, &timezone.Cities, &timezone.UTCOffset,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to scan full user: %w", err)
	}

	res := converter.ToFullUserFromStorage(
		user,
		userParams,
		sex,
		physicalActivity,
		climate,
		timezone,
	)

	return res, nil
}
