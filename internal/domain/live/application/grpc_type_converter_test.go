package application

import (
	"testing"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	mockEntity "github.com/sokorahen-szk/rust-live/tests/domain/live/entity"
	"github.com/stretchr/testify/assert"
)

func Test_ToGrpcLiveVideos(t *testing.T) {
	a := assert.New(t)
	t.Run("0件の場合、空配列を返すこと", func(t *testing.T) {
		emptyLiveVideos := []*entity.LiveVideo{}
		toGrpcLiveVideos := ToGrpcLiveVideos(emptyLiveVideos)
		a.IsType([]*pb.LiveVideo{}, toGrpcLiveVideos)
		a.Len(toGrpcLiveVideos, 0)
	})
	t.Run("2件以上の場合、api.LiveVideo配列を返すこと", func(t *testing.T) {
		liveVideos := []*entity.LiveVideo{
			mockEntity.NewMockLiveVideo(1),
			mockEntity.NewMockLiveVideo(2),
		}

		expectedGrpcLiveVideos := []*pb.LiveVideo{
			{
				Id:               liveVideos[0].Id.Int32(),
				BroadcastId:      liveVideos[0].GetBroadcastId().String(),
				Title:            liveVideos[0].GetTitle().String(),
				Url:              liveVideos[0].GetUrl().String(),
				Stremer:          liveVideos[0].GetStremer().String(),
				Viewer:           liveVideos[0].GetViewer().Int32(),
				VideoPlatform:    pb.VideoPlatform(liveVideos[0].GetPlatform().Int32()),
				ThumbnailImage:   liveVideos[0].GetThumbnailImage().String(),
				StartedDatetime:  liveVideos[0].GetStartedDatetime().RFC3339(),
				ElapsedTimes:     liveVideos[0].GetElapsedTimes().Int32(),
				ElapsedTimesText: liveVideos[0].GetElapsedTimes().Text(),
			},
			{
				Id:               liveVideos[1].Id.Int32(),
				BroadcastId:      liveVideos[1].GetBroadcastId().String(),
				Title:            liveVideos[1].GetTitle().String(),
				Url:              liveVideos[1].GetUrl().String(),
				Stremer:          liveVideos[1].GetStremer().String(),
				Viewer:           liveVideos[1].GetViewer().Int32(),
				VideoPlatform:    pb.VideoPlatform(liveVideos[1].GetPlatform().Int32()),
				ThumbnailImage:   liveVideos[1].GetThumbnailImage().String(),
				StartedDatetime:  liveVideos[1].GetStartedDatetime().RFC3339(),
				ElapsedTimes:     liveVideos[1].GetElapsedTimes().Int32(),
				ElapsedTimesText: liveVideos[1].GetElapsedTimes().Text(),
			},
		}

		toGrpcLiveVideos := ToGrpcLiveVideos(liveVideos)
		a.IsType([]*pb.LiveVideo{}, toGrpcLiveVideos)
		a.Equal(expectedGrpcLiveVideos, toGrpcLiveVideos)
		a.Len(toGrpcLiveVideos, 2)
	})
}

func Test_ToGrpcPagination(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		name string
		arg  *pb.Pagination
	}{
		{
			"リミット=10、レコード=20件、現ページ=1、全ページ=2、前ページ=1、次ページ=2",
			&pb.Pagination{
				Limit:      10,
				Page:       1,
				Prev:       1,
				Next:       2,
				TotalPage:  2,
				TotalCount: 20,
			},
		},
		{
			"リミット=10、レコード=21件、現ページ=2、全ページ=3、前ページ=1、次ページ=3",
			&pb.Pagination{
				Limit:      10,
				Page:       2,
				Prev:       1,
				Next:       3,
				TotalPage:  3,
				TotalCount: 21,
			},
		},
		{
			"リミット=10、レコード=14件、現ページ=14、全ページ=14、前ページ=13、次ページ=14",
			&pb.Pagination{
				Limit:      1,
				Page:       14,
				Prev:       13,
				Next:       14,
				TotalPage:  14,
				TotalCount: 14,
			},
		},
		{
			"リミット=30、レコード=30件、現ページ=1、全ページ=1、前ページ=1、次ページ=1",
			&pb.Pagination{
				Limit:      30,
				Page:       1,
				Prev:       1,
				Next:       1,
				TotalPage:  1,
				TotalCount: 30,
			},
		},
	}
	for _, p := range tests {
		t.Run(p.name, func(t *testing.T) {
			actual := ToGrpcPagination(
				int(p.arg.Page),
				int(p.arg.Limit),
				int(p.arg.TotalCount),
			)
			a.Equal(p.arg, actual)
		})
	}
}
