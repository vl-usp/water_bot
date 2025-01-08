package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/service/refs"
	storageMocks "github.com/vl-usp/water_bot/internal/storage/mocks"
)

func TestSexList(t *testing.T) {
	t.Parallel()

	type refsStorerMockFunc func(mc *minimock.Controller) refs.Storer

	type args struct {
		ctx context.Context
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		storageErr = fmt.Errorf("storage error")

		res = []model.Sex{
			{
				ID:        1,
				Key:       "male",
				Name:      "Мужчина",
				WaterCoef: 1.0,
			},
			{
				ID:        2,
				Key:       "female",
				Name:      "Женщина",
				WaterCoef: 0.8,
			},
		}
	)

	tests := []struct {
		name           string
		args           args
		want           []model.Sex
		err            error
		refsStorerMock refsStorerMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			want: res,
			err:  nil,
			refsStorerMock: func(mc *minimock.Controller) refs.Storer {
				storer := storageMocks.NewReferenceMock(mc)
				storer.SexListMock.Return(res, nil)
				return storer
			},
		},
		{
			name: "error",
			args: args{
				ctx: ctx,
			},
			want: nil,
			err:  storageErr,
			refsStorerMock: func(mc *minimock.Controller) refs.Storer {
				storer := storageMocks.NewReferenceMock(mc)
				storer.SexListMock.Return(nil, storageErr)
				return storer
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			refsStorerMock := tt.refsStorerMock(mc)
			service := refs.New(refsStorerMock)

			res, err := service.SexList(tt.args.ctx)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}

}
