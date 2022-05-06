package controller

import (
	"context"
	"testing"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/stretchr/testify/assert"
)

func Test_LiveController_ListLiveVideos(t *testing.T) {
	t.Skip()
	a := assert.New(t)
	ctx := context.Background()

	t.Run("正常に処理が終了すること", func(t *testing.T) {
		req := &pb.ListLiveVideosRequest{}

		c := LiveController{}

		res, err := c.ListLiveVideos(ctx, req)
		a.NoError(err)
		a.IsType(&pb.ListLiveVideosResponse{}, res)
	})
}
