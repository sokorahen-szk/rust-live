package application_live

import (
	"fmt"
	"math"
	"strconv"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
)

func ToGrpcLiveVideos(liveVideos []*entity.LiveVideo) []*pb.LiveVideo {
	grpcLiveVideos := make([]*pb.LiveVideo, 0)
	for _, liveVideo := range liveVideos {
		grpcLiveVideos = append(grpcLiveVideos, &pb.LiveVideo{
			Id:               liveVideo.GetId().Int32(),
			BroadcastId:      liveVideo.GetBroadcastId().String(),
			Title:            liveVideo.GetTitle().String(),
			Url:              liveVideo.GetUrl().String(),
			Stremer:          liveVideo.GetStremer().String(),
			Viewer:           liveVideo.GetViewer().Int32(),
			VideoPlatform:    pb.VideoPlatform(liveVideo.GetPlatform().Int32()),
			ThumbnailImage:   liveVideo.GetThumbnailImage().String(),
			StartedDatetime:  liveVideo.GetStartedDatetime().RFC3339(),
			ElapsedTimes:     liveVideo.GetElapsedTimes().Int32(),
			ElapsedTimesText: liveVideo.GetElapsedTimes().Text(),
		})
	}
	return grpcLiveVideos
}

func ToGrpcPagination(page int, limit int, total int) *pb.Pagination {
	f := math.Ceil(float64(total) / float64(limit))
	totalPage, err := strconv.Atoi(fmt.Sprintf("%.0f", f))
	if err != nil {
		panic(err)
	}

	prev := page
	if page > 1 {
		prev--
	}

	next := page
	if next < totalPage {
		next++
	}

	return &pb.Pagination{
		Limit:      int32(limit),
		Page:       int32(page),
		Prev:       int32(prev),
		Next:       int32(next),
		TotalPage:  int32(totalPage),
		TotalCount: int32(total),
	}
}
