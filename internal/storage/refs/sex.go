package refs

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/storage/refs/converter"
	repoModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
	"github.com/vl-usp/water_bot/pkg/client/db"
)

// SexList returns a list of sexes
func (s *store) SexList(ctx context.Context) ([]model.Sex, error) {
	builder := sq.Select(idColumn, keyColumn, nameColumn, waterCoefColumn).
		PlaceholderFormat(sq.Dollar).
		From(sexTable).
		OrderBy(idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.SexList",
		QueryRaw: query,
	}

	sexList := make([]repoModel.Sex, 0)
	err = s.db.DB().ScanAllContext(ctx, &sexList, q, args...)
	if err != nil {
		return nil, err
	}

	sexListModels := make([]model.Sex, 0, len(sexList))
	for _, sex := range sexList {
		sexListModels = append(sexListModels, converter.ToSexFromRepo(sex))
	}

	return sexListModels, nil
}
