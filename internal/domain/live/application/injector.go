package application

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
)

func NewInjectListLiveVideosUsecase(ctx context.Context) list.ListLiveVideosUsecaseInterface {
	config := cfg.NewConfig()
	redis := redis.NewRedis(ctx, config)

	var liveVideoRepository repository.LiveVideoRepositoryInterface
	if config.IsProd() {
		liveVideoRepository = live.NewLiveVideoRepository(redis)
	} else {
		liveVideoRepository = live.NewMockLiveVideoRepository()
	}

	return NewListLiveVideosUsecase(liveVideoRepository)
}
