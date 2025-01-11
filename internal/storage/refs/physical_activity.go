package refs

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/storage/refs/converter"
	repoModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
	"github.com/vl-usp/water_bot/pkg/client/db"
)

// PhysicalActivityList returns a list of physical activities
func (s *store) PhysicalActivityList(ctx context.Context) ([]model.PhysicalActivity, error) {
	builder := sq.Select(idColumn, keyColumn, nameColumn, waterCoefColumn).
		PlaceholderFormat(sq.Dollar).
		From(physicalActivityTable).
		OrderBy(idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build physical activity list query: %w", err)
	}

	q := db.Query{
		Name:     "refs_storage.PhysicalActivityList",
		QueryRaw: query,
	}

	physicalActivityList := make([]repoModel.PhysicalActivity, 0)
	err = s.db.DB().ScanAllContext(ctx, &physicalActivityList, q, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get physical activity list: %w", err)
	}

	physicalActivityListModels := make([]model.PhysicalActivity, 0, len(physicalActivityList))
	for _, physicalActivity := range physicalActivityList {
		physicalActivityListModels = append(physicalActivityListModels, *converter.ToPhysicalActivityFromStorage(physicalActivity))
	}

	return physicalActivityListModels, nil
}
