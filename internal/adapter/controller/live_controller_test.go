package controller

import (
	"context"
	"testing"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/stretchr/testify/assert"
)

func Test_LiveController_ListLiveVideos(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	t.Run("正常に処理が終了すること", func(t *testing.T) {
		req := &pb.ListLiveVideosRequest{
			Page: 1,
		}

		c := LiveController{}

		res, err := c.ListLiveVideos(ctx, req)
		a.NoError(err)
		a.NotNil(res)
		a.IsType(&pb.ListLiveVideosResponse{}, res)
	})
	t.Run("Usecaseでエラーが発生した場合、異常終了すること", func(t *testing.T) {
		ctxWithError := context.WithValue(ctx, "test", "list_live_videos_usecase_error")

		req := &pb.ListLiveVideosRequest{
			Page: 1,
		}

		c := LiveController{}

		res, err := c.ListLiveVideos(ctxWithError, req)
		a.Error(err)
		a.Nil(res)
		a.Equal("list live videos usecase error", err.Error())
	})
}
