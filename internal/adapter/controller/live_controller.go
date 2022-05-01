package controller

import (
	"context"

	pb "github.com/sokorahen-szk/rust-live/api/proto"

	"github.com/sokorahen-szk/rust-live/internal/adapter/controller/form"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/application"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
)

type LiveController struct {
	pb.UnimplementedLiveServiceServer
}

func (s *LiveController) ListLiveVideos(ctx context.Context, req *pb.ListLiveVideosRequest) (*pb.ListLiveVideosResponse, error) {
	formData := form.NewListLiveVideosForm(req)
	err := form.Validate(formData)
	if err != nil {
		return nil, err
	}

	usecase := application.NewInjectListLiveVideosUsecase(ctx)
	return usecase.Handle(ctx, list.NewListLiveVideosInput(formData.GetSearchKeywords()))
}
