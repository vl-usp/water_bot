package refs

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// ClimateList returns a list of climates
func (s *serv) ClimateList(ctx context.Context) ([]model.Climate, error) {
	return s.storage.ClimateList(ctx)
}
