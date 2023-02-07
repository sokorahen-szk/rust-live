package application_live

import (
	"context"
	"errors"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
	mockEntity "github.com/sokorahen-szk/rust-live/tests/domain/live/entity"
)

type listLiveVideosUsecaseMock struct{}

func NewListLiveVideosUsecaseMock() list.ListLiveVideosUsecaseInterface {
	return listLiveVideosUsecaseMock{}
}

func (ins listLiveVideosUsecaseMock) Handle(ctx context.Context, input *list.ListLiveVideoInput) (*pb.ListLiveVideosResponse, error) {
	if ctx.Value("test") == "list_live_videos_usecase_error" {
		return nil, errors.New("list live videos usecase error")
	}

	liveVideos := []*entity.LiveVideo{
		mockEntity.NewMockLiveVideo(1),
	}

	res := &pb.ListLiveVideosResponse{
		LiveVideos: ToGrpcLiveVideos(liveVideos),
		Pagination: ToGrpcPagination(input.Page(), input.Limit(), 1),
	}

	return res, nil
}
