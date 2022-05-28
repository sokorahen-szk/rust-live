package controller

import (
	"context"

	pb "github.com/sokorahen-szk/rust-live/api/proto"

	"github.com/sokorahen-szk/rust-live/internal/adapter/controller/form"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/application"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
)

type LiveController struct {
	pb.UnimplementedLiveServiceServer
}

const (
	listLiveVideoRequestDefaultLimit int = 10
)

func (s *LiveController) ListLiveVideos(ctx context.Context, req *pb.ListLiveVideosRequest) (*pb.ListLiveVideosResponse, error) {
	formData := form.NewListLiveVideosForm(req)
	err := form.Validate(formData)
	if err != nil {
		return nil, err
	}

	limit := formData.GetLimit()
	if limit == 0 {
		limit = listLiveVideoRequestDefaultLimit
	}

	sortKey := entity.NewLiveVideoSortKeyFromInt(formData.GetSort())

	platforms := make([]*entity.Platform, len(formData.GetPlatforms()))
	for _, platform := range formData.GetPlatforms() {
		platforms = append(platforms, entity.NewPlatformFromInt(platform))
	}
	input := list.NewListLiveVideoInput(
		formData.GetSearchKeywords(),
		platforms,
		sortKey,
		formData.GetPage(),
		limit,
	)

	usecase := application.NewInjectListLiveVideosUsecase(ctx)
	return usecase.Handle(ctx, input)
}
