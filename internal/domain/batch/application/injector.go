package application

import (
	"context"
	"net/http"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	usecaseBatch "github.com/sokorahen-szk/rust-live/internal/usecase/batch"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch/twitch"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
)

func NewInjectFetchLiveVideosUsecase(ctx context.Context) usecaseBatch.FetchLiveVideosUsecaseInterface {
	config := cfg.NewConfig()
	redis := redis.NewRedis(ctx, config)

	var liveVideoRepository repository.LiveVideoRepositoryInterface
	if config.IsProd() {
		liveVideoRepository = live.NewLiveVideoRepository(redis)
	} else {
		liveVideoRepository = live.NewMockLiveVideoRepository()
	}

	client := batch.NewHttpClient(http.MethodGet, nil)
	twitchApiClient := twitch.NewTwitchApiClient(client, config)

	return NewFetchLiveVideosUsecase(liveVideoRepository, twitchApiClient)
}
