package application

import (
	"context"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
)

type listLiveVideosUsecase struct {
	liveVideoRepository repository.LiveVideoRepositoryInterface
}

func NewListLiveVideosUsecase(
	liveVideoRepository repository.LiveVideoRepositoryInterface,
) list.ListLiveVideosUsecaseInterface {
	return listLiveVideosUsecase{
		liveVideoRepository: liveVideoRepository,
	}
}

func (ins listLiveVideosUsecase) Handle(ctx context.Context, input *list.ListLiveVideoInput) (*pb.ListLiveVideosResponse, error) {
	liveVideos, err := ins.liveVideoRepository.List(ctx, input)
	if err != nil {
		return nil, err
	}

	liveVideoTotalCount, err := ins.liveVideoRepository.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ListLiveVideosResponse{
		LiveVideos: ToGrpcLiveVideos(liveVideos),
		Pagination: ToGrpcPagination(input.Page(), input.Limit(), liveVideoTotalCount),
	}, nil
}
