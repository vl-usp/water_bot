package user

import (
	"github.com/vl-usp/water_bot/internal/client/db"
	"github.com/vl-usp/water_bot/internal/repository"
	"github.com/vl-usp/water_bot/internal/service"
)

type serv struct {
	userRepo  repository.UserRepository
	txManager db.TxManager
}

// NewService creates a new user service.
func NewService(
	userRepo repository.UserRepository,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		userRepo:  userRepo,
		txManager: txManager,
	}
}

// NewMockService creates a new mock user service.
func NewMockService(deps ...interface{}) service.UserService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.UserRepository:
			srv.userRepo = s
		case db.TxManager:
			srv.txManager = s
		}
	}

	return &srv
}
