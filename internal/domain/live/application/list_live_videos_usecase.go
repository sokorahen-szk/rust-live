package application

import (
	"context"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
)

type ListLiveVideosUsecase struct {
	liveVideoRepository repository.LiveVideoRepositoryInterface
}

func NewListLiveVideosUsecase(
	liveVideoRepository repository.LiveVideoRepositoryInterface,
) list.ListLiveVideosUsecaseInterface {
	return ListLiveVideosUsecase{
		liveVideoRepository: liveVideoRepository,
	}
}

func (ins ListLiveVideosUsecase) Handle(ctx context.Context, input *list.ListLiveVideosInput) (*pb.ListLiveVideosResponse, error) {
	liveVideos, err := ins.liveVideoRepository.List(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.ListLiveVideosResponse{
		LiveVideos: ToGrpcLiveVideos(liveVideos),
	}, nil
}
