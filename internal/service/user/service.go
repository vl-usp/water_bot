package user

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/service"
	"github.com/vl-usp/water_bot/pkg/client/db"
)

// Storer represents a storage for user data.
type Storer interface {
	CreateUser(ctx context.Context, user model.User) (int64, error)
	UpdateUser(ctx context.Context, id int64, user model.User) error
	GetUser(ctx context.Context, id int64) (*model.User, error)

	CreateUserParams(ctx context.Context, params model.UserParams) (int64, error)
	UpdateUserParams(ctx context.Context, id int64, params model.UserParams) error
	GetFullUserParams(ctx context.Context, params model.UserParams) (*model.UserParams, error)
}

// Cacher represents a cache for user params, that needed for temporary storage of user data.
type Cacher interface {
	SaveUserParam(ctx context.Context, userID int64, field string, value interface{}) error
	GetUserParams(ctx context.Context, userID int64) (*model.UserParams, error)
}

type serv struct {
	userStore Storer
	userCache Cacher
	txManager db.TxManager
}

// New creates a new user service.
func New(
	userStore Storer,
	userCache Cacher,
	txManager db.TxManager,
) service.User {
	return &serv{
		userStore: userStore,
		userCache: userCache,
		txManager: txManager,
	}
}

// NewMockService creates a new mock user service.
func NewMockService(deps ...interface{}) service.User {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case Storer:
			srv.userStore = s
		case Cacher:
			srv.userCache = s
		case db.TxManager:
			srv.txManager = s
		}
	}

	return &srv
}
