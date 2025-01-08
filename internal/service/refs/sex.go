package refs

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// SexList returns a list of sexes
func (s *serv) SexList(ctx context.Context) ([]model.Sex, error) {
	return s.storage.SexList(ctx)
}
