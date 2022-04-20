package controller

import (
	"context"

	pb "github.com/sokorahen-szk/rust-live/api/proto"

	"github.com/sokorahen-szk/rust-live/internal/adapter/controller/form"
)

type LiveController struct {
	pb.UnimplementedLiveServiceServer
}

func (s *LiveController) ListLiveVideos(ctx context.Context, req *pb.ListLiveVideosRequest) (*pb.ListLiveVideosResponse, error) {
	err := form.Validate(form.NewListLiveVideosForm(req))
	if err != nil {
		return nil, err
	}

	return &pb.ListLiveVideosResponse{}, nil
}
