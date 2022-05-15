package application

import (
	"context"
	"testing"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
	"github.com/stretchr/testify/assert"
)

func Test_ListLiveVideosUsecase_Handle(t *testing.T) {
	t.Skip()
	a := assert.New(t)
	ctx := context.Background()

	listLiveUsecase := NewInjectListLiveVideosUsecase(ctx)

	t.Run("正常系/配信動画リスト取得", func(t *testing.T) {
		tests := []struct {
			name string
			arg  string
			want int
		}{
			{"検索キーワードが空の場合、10件を返す.", "", 10},
			{"検索キーワードを指定して、一致するデータがない場合、0件を返す.", "太郎", 0},
			{"検索キーワードを指定して、一致するデータが1件の場合、1件を返す.", "stremer2", 1},
		}
		for _, p := range tests {

			input := list.NewListLiveVideoInput(p.arg)

			t.Run(p.name, func(t *testing.T) {
				res, err := listLiveUsecase.Handle(ctx, input)
				a.IsType(res, &pb.ListLiveVideosResponse{})
				a.Len(res.LiveVideos, p.want)
				a.NoError(err)
			})
		}

	})
	t.Run("異常系/配信動画リスト取得失敗", func(t *testing.T) {
		searchKeywords := "keywords"
		ctxWithError := context.WithValue(ctx, "error", "error")

		listLiveUsecase := NewInjectListLiveVideosUsecase(ctx)
		input := list.NewListLiveVideoInput(searchKeywords)

		res, err := listLiveUsecase.Handle(ctxWithError, input)
		a.Nil(res)
		a.Error(err)
	})
}
