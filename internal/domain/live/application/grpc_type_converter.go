package application

import (
	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
)

func ToGrpcLiveVideos(liveVideos []*entity.LiveVideo) []*pb.LiveVideo {
	grpcLiveVideos := make([]*pb.LiveVideo, 0)
	for _, liveVideo := range liveVideos {
		grpcLiveVideos = append(grpcLiveVideos, &pb.LiveVideo{
			Id:              liveVideo.GetId().Int32(),
			Title:           liveVideo.GetTitle().String(),
			Stremer:         liveVideo.GetStremer().String(),
			Viewer:          liveVideo.GetViewer().Int32(),
			ThumbnailImage:  liveVideo.GetThumbnailImage().String(),
			StartedDatetime: liveVideo.GetStartedDatetime().RFC3339(),
			ElapsedTimes:    liveVideo.GetElapsedTimes().Int32(),
		})
	}
	return grpcLiveVideos
}
