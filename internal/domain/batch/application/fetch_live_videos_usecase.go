package application

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/usecase/batch"
)

type fetchLiveVideosUsecase struct {
	liveVideoRepository repository.LiveVideoRepositoryInterface
}

func NewFetchLiveVideosUsecase(
	liveVideoRepository repository.LiveVideoRepositoryInterface,
) batch.FetchLiveVideosUsecaseInterface {
	return fetchLiveVideosUsecase{
		liveVideoRepository: liveVideoRepository,
	}
}

func (usecase fetchLiveVideosUsecase) Handle(ctx context.Context) error {
	// TODO:
	return nil
}
