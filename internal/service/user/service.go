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

func NewService(
	userRepo repository.UserRepository,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		userRepo:  userRepo,
		txManager: txManager,
	}
}
