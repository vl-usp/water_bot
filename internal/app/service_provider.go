package app

import (
	"context"

	"github.com/vl-usp/water_bot/internal/service"
	refsService "github.com/vl-usp/water_bot/internal/service/refs"
	userService "github.com/vl-usp/water_bot/internal/service/user"
)

type serviceProvider struct {
	storageProvider *storageProvider

	userService service.User
	refsService service.Reference
}

func newServiceProvider(storageProvider *storageProvider) *serviceProvider {
	return &serviceProvider{
		storageProvider: storageProvider,
	}
}

// UserService returns a user service.
func (s *serviceProvider) UserService(ctx context.Context) service.User {
	if s.userService == nil {
		s.userService = userService.New(
			s.storageProvider.UserStorage(ctx),
			s.storageProvider.UserCache(ctx),
			s.storageProvider.TxManager(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) ReferenceService(ctx context.Context) service.Reference {
	if s.refsService == nil {
		s.refsService = refsService.New(
			s.storageProvider.ReferenceStorage(ctx),
		)
	}

	return s.refsService
}
