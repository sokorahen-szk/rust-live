package application_live

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
)

func NewInjectListLiveVideosUsecase(ctx context.Context) list.ListLiveVideosUsecaseInterface {
	config := cfg.NewConfig()
	redis := redis.NewRedis(ctx, config)

	liveVideoRepository := live.NewLiveVideoRepository(redis)
	if config.IsProd() || config.IsDev() {
		return NewListLiveVideosUsecase(liveVideoRepository)
	}

	return NewListLiveVideosUsecaseMock()
}
