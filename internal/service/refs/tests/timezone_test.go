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

func TestTimezoneList(t *testing.T) {
	t.Parallel()

	type refsStorerMockFunc func(mc *minimock.Controller) refs.Storer

	type args struct {
		ctx context.Context
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		storageErr = fmt.Errorf("storage error")

		res = []model.Timezone{
			{
				ID:        1,
				Name:      "UTC-12",
				Cities:    "Moscow, Saint Petersburg, Novosibirsk",
				UTCOffset: -12,
			},
			{
				ID:        2,
				Name:      "UTC-11",
				Cities:    "Anchorage, Juneau, Metlakatla",
				UTCOffset: -11,
			},
			{
				ID:        3,
				Name:      "UTC-10",
				Cities:    "Hawaii",
				UTCOffset: -10,
			},
			{
				ID:        4,
				Name:      "UTC-9",
				Cities:    "Alaska",
				UTCOffset: -9,
			},
		}
	)

	tests := []struct {
		name           string
		args           args
		want           []model.Timezone
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
				storer.TimezoneListMock.Return(res, nil)
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
				storer.TimezoneListMock.Return(nil, storageErr)
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

			res, err := service.TimezoneList(tt.args.ctx)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
