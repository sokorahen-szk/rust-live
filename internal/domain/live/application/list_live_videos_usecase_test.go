package application

import (
	"context"
	"testing"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
	"github.com/stretchr/testify/assert"
)

func Test_NewListLiveVideosUsecase(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	searchKeywords := "test"

	t.Run("正常系/配信動画リスト取得が正しくできること", func(t *testing.T) {
		listLiveUsecase := NewInjectListLiveVideosUsecase()
		input := list.NewListLiveVideosInput(searchKeywords)

		res, err := listLiveUsecase.Handle(ctx, input)
		a.IsType(res, &pb.ListLiveVideosResponse{})
		a.Len(res.LiveVideos, 10)
		a.NoError(err)
	})
	t.Run("異常系/配信動画リスト取得失敗", func(t *testing.T) {
		ctxWithError := context.WithValue(ctx, "error", "error")

		listLiveUsecase := NewInjectListLiveVideosUsecase()
		input := list.NewListLiveVideosInput(searchKeywords)

		res, err := listLiveUsecase.Handle(ctxWithError, input)
		a.Nil(res)
		a.Error(err)
	})
}
