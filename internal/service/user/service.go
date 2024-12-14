package user

import (
	"github.com/vl-usp/water_bot/internal/repository"
	"github.com/vl-usp/water_bot/internal/service"
	"github.com/vl-usp/water_bot/pkg/client/db"
)

type serv struct {
	userRepo     repository.UserRepository
	userDataRepo repository.UserDataRepository
	txManager    db.TxManager
}

// NewService creates a new user service.
func NewService(
	userRepo repository.UserRepository,
	userDataRepo repository.UserDataRepository,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		userRepo:     userRepo,
		userDataRepo: userDataRepo,
		txManager:    txManager,
	}
}

// NewMockService creates a new mock user service.
func NewMockService(deps ...interface{}) service.UserService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.UserRepository:
			srv.userRepo = s
		case repository.UserDataRepository:
			srv.userDataRepo = s
		case db.TxManager:
			srv.txManager = s
		}
	}

	return &srv
}
