package application

import (
	"context"
	"net/http"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	usecaseBatch "github.com/sokorahen-szk/rust-live/internal/usecase/batch"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch/twitch"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
	postgresqlLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql/live"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	redisLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
)

func NewInjectFetchLiveVideosUsecase(ctx context.Context) usecaseBatch.FetchLiveVideosUsecaseInterface {
	config := cfg.NewConfig()
	redis := redis.NewRedis(ctx, config)
	postgresql := postgresql.NewPostgreSQL(config)

	var liveVideoRepository repository.LiveVideoRepositoryInterface
	if config.IsProd() {
		liveVideoRepository = redisLive.NewLiveVideoRepository(redis)
	} else {
		liveVideoRepository = redisLive.NewMockLiveVideoRepository()
	}

	client := batch.NewHttpClient(http.MethodGet, nil)
	archiveVideoRepository := postgresqlLive.NewArchiveVideoRepository(postgresql)
	twitchApiClient := twitch.NewTwitchApiClient(client, config)

	return NewFetchLiveVideosUsecase(
		liveVideoRepository,
		archiveVideoRepository,
		twitchApiClient,
	)
}
