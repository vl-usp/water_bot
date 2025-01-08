package refs

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/service"
)

// Storer represents a reference service.
type Storer interface {
	ClimateList(ctx context.Context) ([]model.Climate, error)
	PhysicalActivityList(ctx context.Context) ([]model.PhysicalActivity, error)
	SexList(ctx context.Context) ([]model.Sex, error)
	TimezoneList(ctx context.Context) ([]model.Timezone, error)
}

type serv struct {
	storage Storer
}

// New creates a new user service.
func New(
	storage Storer,
) service.Reference {
	return &serv{
		storage: storage,
	}
}
