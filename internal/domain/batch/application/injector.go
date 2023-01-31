package application

import (
	"context"
	"time"

	usecaseBatch "github.com/sokorahen-szk/rust-live/internal/usecase/batch"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch/twitch"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
	postgresqlLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql/live"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	redisLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
	httpClient "github.com/sokorahen-szk/rust-live/pkg/http"
)

func NewInjectFetchLiveVideosUsecase(ctx context.Context) usecaseBatch.FetchLiveVideosUsecaseInterface {
	config := cfg.NewConfig()
	redis := redis.NewRedis(ctx, config)
	postgresql := postgresql.NewPostgreSQL(config)

	client := httpClient.NewHttpClient(nil)
	archiveVideoRepository := postgresqlLive.NewArchiveVideoRepository(postgresql)
	liveVideoRepository := redisLive.NewLiveVideoRepository(redis)
	twitchApiClient := twitch.NewTwitchApiClient(client, config)

	return NewFetchLiveVideosUsecase(
		liveVideoRepository,
		archiveVideoRepository,
		twitchApiClient,
		time.Now,
	)
}

func NewInjectUpdateLiveVideosUsecase(ctx context.Context) usecaseBatch.UpdateLiveVideosUsecaseInterface {
	config := cfg.NewConfig()
	postgresql := postgresql.NewPostgreSQL(config)

	archiveVideoRepository := postgresqlLive.NewArchiveVideoRepository(postgresql)

	client := httpClient.NewHttpClient(nil)
	twitchApiClient := twitch.NewTwitchApiClient(client, config)

	return NewUpdateLiveVideosUsecase(
		archiveVideoRepository,
		twitchApiClient,
		time.Now,
	)
}
