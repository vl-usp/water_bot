package refs

import (
	"context"

	"github.com/vl-usp/water_bot/internal/model"
)

// TimezoneList returns a list of timezones
func (s *serv) TimezoneList(ctx context.Context) ([]model.Timezone, error) {
	return s.storage.TimezoneList(ctx)
}
