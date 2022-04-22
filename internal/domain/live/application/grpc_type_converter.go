package application

import (
	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
)

func ToGrpcLiveVideos(liveVideos []*entity.LiveVideo) []*pb.LiveVideo {
	grpcLiveVideos := make([]*pb.LiveVideo, 0)
	for _, liveVideo := range liveVideos {
		grpcLiveVideos = append(grpcLiveVideos, &pb.LiveVideo{
			Id:              liveVideo.Id().Int32(),
			Title:           liveVideo.Title().String(),
			Stremer:         liveVideo.Stremer().String(),
			ThumbnailImage:  liveVideo.ThumbnailImage().String(),
			StartedDatetime: liveVideo.StartedDatetime().RFC3339(),
			ElapsedTimes:    liveVideo.ElapsedTimes().Int32(),
		})
	}
	return grpcLiveVideos
}
