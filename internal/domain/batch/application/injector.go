package application

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/usecase/batch"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
)

func NewInjectFetchLiveVideosUsecase(ctx context.Context) batch.FetchLiveVideosUsecaseInterface {
	config := cfg.NewConfig()
	redis := redis.NewRedis(ctx, config)

	var liveVideoRepository repository.LiveVideoRepositoryInterface
	if config.IsProd() {
		liveVideoRepository = live.NewLiveVideoRepository(redis)
	} else {
		liveVideoRepository = live.NewMockLiveVideoRepository()
	}

	return NewFetchLiveVideosUsecase(liveVideoRepository)
}
