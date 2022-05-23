package application

import (
	"context"
	"testing"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
	"github.com/stretchr/testify/assert"
)

func Test_ListLiveVideosUsecase_Handle(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	searchKeywords := "keywords"
	platforms := []*entity.Platform{}
	sortKey := entity.NewLiveVideoSortKey(0)

	t.Run("異常系/配信動画リスト取得失敗", func(t *testing.T) {
		ctxWithError := context.WithValue(ctx, "test", "error")

		listLiveUsecase := NewInjectListLiveVideosUsecase(ctx)
		input := list.NewListLiveVideoInput(searchKeywords, platforms, sortKey, 1, 10)

		res, err := listLiveUsecase.Handle(ctxWithError, input)
		a.Nil(res)
		a.Error(err)
	})
}
