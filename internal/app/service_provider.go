package app

import (
	"context"

	"github.com/vl-usp/water_bot/internal/client/db"
	"github.com/vl-usp/water_bot/internal/client/db/pg"
	"github.com/vl-usp/water_bot/internal/client/db/transaction"
	"github.com/vl-usp/water_bot/internal/closer"
	"github.com/vl-usp/water_bot/internal/config"
	"github.com/vl-usp/water_bot/internal/repository"
	userRepository "github.com/vl-usp/water_bot/internal/repository/user"
	"github.com/vl-usp/water_bot/internal/service"
	userService "github.com/vl-usp/water_bot/internal/service/user"
	"github.com/vl-usp/water_bot/pkg/logger"
)

type serviceProvider struct {
	systemConfig config.SystemConfig
	pgConfig     config.PGConfig
	tgConfig     config.TGConfig

	dbClient       db.Client
	txManager      db.TxManager
	userRepository repository.UserRepository
	userService    service.UserService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// SystemConfig returns a config that stores system settings.
func (s *serviceProvider) SystemConfig() config.SystemConfig {
	if s.systemConfig == nil {
		cfg, err := config.NewSystemConfig()
		if err != nil {
			logger.Get("app", "s.SystemConfig").Error("failed to get system config", "error", err.Error())
		}

		s.systemConfig = cfg
	}

	return s.systemConfig
}

// PGConfig returns a pg config.
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			logger.Get("app", "s.PGConfig").Error("failed to get pg config", "error", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// TGConfig returns a tg config.
func (s *serviceProvider) TGConfig() config.TGConfig {
	if s.tgConfig == nil {
		cfg, err := config.NewTGConfig()
		if err != nil {
			logger.Get("app", "s.TGConfig").Error("failed to get tg config", "error", err.Error())
		}

		s.tgConfig = cfg
	}

	return s.tgConfig
}

// DBClient returns a db client.
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN(), s.SystemConfig().Debug())
		if err != nil {
			logger.Get("app", "s.DBClient").Error("failed to create db client", "error", err.Error())
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			logger.Get("app", "s.DBClient").Error("ping error", "error", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager returns a transaction manager.
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// UserRepository returns a user repository.
func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

// UserService returns a user service.
func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}
