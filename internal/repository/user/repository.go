package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/vl-usp/water_bot/internal/client/db"
	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/repository"
	"github.com/vl-usp/water_bot/internal/repository/user/converter"
	modelRepo "github.com/vl-usp/water_bot/internal/repository/user/model"
)

const (
	tableName = "users"

	idColumn        = "id"
	firstNameColumn = "first_name"
	lastNameColumn  = "last_name"
	usernameColumn  = "username"
	langCodeColumn  = "language_code"
	createdAtColumn = "created_at"
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

// Create creates a new user.
func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(idColumn, firstNameColumn, lastNameColumn, usernameColumn, langCodeColumn, createdAtColumn).
		Values(user.ID, user.FirstName, user.LastName, user.Username, user.LanguageCode, user.CreatedAt).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get returns a user by id.
func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, firstNameColumn, lastNameColumn, usernameColumn, langCodeColumn, createdAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.LanguageCode,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
