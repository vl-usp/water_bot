package app

import (
	"context"

	"github.com/vl-usp/water_bot/internal/closer"
	"github.com/vl-usp/water_bot/internal/storage"
	refsStorage "github.com/vl-usp/water_bot/internal/storage/refs"
	userStorage "github.com/vl-usp/water_bot/internal/storage/user"
	userCacheStorage "github.com/vl-usp/water_bot/internal/storage/user_cache"
	"github.com/vl-usp/water_bot/pkg/client/cache"
	"github.com/vl-usp/water_bot/pkg/client/cache/redis"
	"github.com/vl-usp/water_bot/pkg/client/db"
	"github.com/vl-usp/water_bot/pkg/client/db/pg"
	"github.com/vl-usp/water_bot/pkg/client/db/transaction"
	"github.com/vl-usp/water_bot/pkg/logger"
)

type storageProvider struct {
	configProvider *configProvider

	dbClient    db.Client
	txManager   db.TxManager
	redisClient cache.Client

	userStorage storage.User
	userCache   storage.UserCache
	refsStorage storage.Reference
}

func newStorageProvider(configProvider *configProvider) *storageProvider {
	return &storageProvider{
		configProvider: configProvider,
	}
}

// DBClient returns a db client.
func (s *storageProvider) DBClient(ctx context.Context) db.Client {
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
func (s *storageProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// RedisClient returns a redis client.
func (s *storageProvider) RedisClient(_ context.Context) cache.Client {
	if s.redisClient == nil {
		s.redisClient = redis.New(s.configProvider.RedisConfig(), logger.Get("pkg/client/cache/redis", ""))
	}

	return s.redisClient
}

// UserStorage returns a user repository.
func (s *storageProvider) UserStorage(ctx context.Context) storage.User {
	if s.userStorage == nil {
		s.userStorage = userStorage.New(s.DBClient(ctx))
	}

	return s.userStorage
}

// UserCache returns a user repository.
func (s *storageProvider) UserCache(ctx context.Context) storage.UserCache {
	if s.userCache == nil {
		s.userCache = userCacheStorage.New(s.RedisClient(ctx))
	}

	return s.userCache
}

// ReferenceStorage returns a repository that store reference data.
func (s *storageProvider) ReferenceStorage(ctx context.Context) storage.Reference {
	if s.refsStorage == nil {
		s.refsStorage = refsStorage.New(s.DBClient(ctx))
	}

	return s.refsStorage
}
