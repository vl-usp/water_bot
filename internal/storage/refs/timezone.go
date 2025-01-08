package refs

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/storage/refs/converter"
	repoModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
	"github.com/vl-usp/water_bot/pkg/client/db"
)

// TimezoneList returns a list of all timezones.
func (s *store) TimezoneList(ctx context.Context) ([]model.Timezone, error) {
	builder := sq.Select(idColumn, nameColumn, citiesColumn, utcOffsetColumn).
		PlaceholderFormat(sq.Dollar).
		From(timezoneTable).
		OrderBy(idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.TimezoneList",
		QueryRaw: query,
	}

	timezoneList := make([]repoModel.Timezone, 0)
	err = s.db.DB().ScanAllContext(ctx, &timezoneList, q, args...)
	if err != nil {
		return nil, err
	}

	timezoneListModels := make([]model.Timezone, 0, len(timezoneList))
	for _, timezone := range timezoneList {
		timezoneListModels = append(timezoneListModels, converter.ToTimezoneFromRepo(timezone))
	}

	return timezoneListModels, nil
}
