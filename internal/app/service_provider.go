package app

import (
	"context"

	"github.com/vl-usp/water_bot/internal/closer"
	"github.com/vl-usp/water_bot/internal/repository"
	userRepository "github.com/vl-usp/water_bot/internal/repository/user"
	userDataRepository "github.com/vl-usp/water_bot/internal/repository/user_data"
	"github.com/vl-usp/water_bot/internal/service"
	userService "github.com/vl-usp/water_bot/internal/service/user"
	"github.com/vl-usp/water_bot/pkg/client/cache"
	"github.com/vl-usp/water_bot/pkg/client/cache/redis"
	"github.com/vl-usp/water_bot/pkg/client/db"
	"github.com/vl-usp/water_bot/pkg/client/db/pg"
	"github.com/vl-usp/water_bot/pkg/client/db/transaction"
	"github.com/vl-usp/water_bot/pkg/logger"
)

type serviceProvider struct {
	configProvider *configProvider

	dbClient           db.Client
	txManager          db.TxManager
	redisClient        cache.Client
	userRepository     repository.UserRepository
	userDataRepository repository.UserDataRepository
	userService        service.UserService
}

func newServiceProvider(configProvider *configProvider) *serviceProvider {
	return &serviceProvider{
		configProvider: configProvider,
	}
}

// DBClient returns a db client.
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(
			ctx,
			s.configProvider.PGConfig().DSN(),
			logger.Get("pkg/client/db", ""),
			s.configProvider.SystemConfig().Debug(),
		)
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

// RedisClient returns a redis client.
func (s *serviceProvider) RedisClient(_ context.Context) cache.Client {
	if s.redisClient == nil {
		s.redisClient = redis.New(s.configProvider.RedisConfig(), logger.Get("pkg/client/cache/redis", ""))
	}

	return s.redisClient
}

// UserRepository returns a user repository.
func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserDataRepository(ctx context.Context) repository.UserDataRepository {
	if s.userDataRepository == nil {
		s.userDataRepository = userDataRepository.NewRepository(s.DBClient(ctx), s.RedisClient(ctx))
	}

	return s.userDataRepository
}

// UserService returns a user service.
func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.UserDataRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}
