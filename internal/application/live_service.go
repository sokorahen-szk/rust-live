package application

import (
	"context"

	pb "github.com/sokorahen-szk/rust-live/api/proto"

	"github.com/sokorahen-szk/rust-live/internal/application/form"
)

type LiveService struct {
	pb.UnimplementedLiveServiceServer
}

func (s *LiveService) ListLiveVideos(ctx context.Context, req *pb.ListLiveVideosRequest) (*pb.ListLiveVideosResponse, error) {
	err := form.Validate(form.NewListLiveVideosForm(req))
	if err != nil {
		return nil, err
	}

	return &pb.ListLiveVideosResponse{}, nil
}
