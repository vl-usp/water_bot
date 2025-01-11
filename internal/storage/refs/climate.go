package refs

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/vl-usp/water_bot/internal/model"
	repoModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
	"github.com/vl-usp/water_bot/pkg/client/db"

	"github.com/vl-usp/water_bot/internal/storage/refs/converter"
)

// ClimateList returns a list of climates
func (s *store) ClimateList(ctx context.Context) ([]model.Climate, error) {
	builder := sq.Select(idColumn, keyColumn, nameColumn, waterCoefColumn).
		PlaceholderFormat(sq.Dollar).
		From(climateTable).
		OrderBy(idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build climate list query: %w", err)
	}

	q := db.Query{
		Name:     "refs_storage.ClimateList",
		QueryRaw: query,
	}

	climateList := make([]repoModel.Climate, 0)
	err = s.db.DB().ScanAllContext(ctx, &climateList, q, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get climate list: %w", err)
	}

	cliamteListModels := make([]model.Climate, 0, len(climateList))
	for _, climate := range climateList {
		cliamteListModels = append(cliamteListModels, *converter.ToClimateFromStorage(climate))
	}

	return cliamteListModels, nil
}
