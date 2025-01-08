package refs

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// PhysicalActivityList returns a list of physical activities
func (s *serv) PhysicalActivityList(ctx context.Context) ([]model.PhysicalActivity, error) {
	return s.storage.PhysicalActivityList(ctx)
}
