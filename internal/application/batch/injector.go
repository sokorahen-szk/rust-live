package application_batch

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

func NewInjectTwitchFetchLiveVideosUsecase(ctx context.Context) usecaseBatch.TwitchFetchLiveVideosUsecaseInterface {
	config := cfg.NewConfig()
	redis := redis.NewRedis(ctx, config)
	postgresql := postgresql.NewPostgreSQL(config)

	client := httpClient.NewHttpClient(nil)
	archiveVideoRepository := postgresqlLive.NewArchiveVideoRepository(postgresql)
	liveVideoRepository := redisLive.NewLiveVideoRepository(redis)
	twitchApiClient := twitch.NewTwitchApiClient(client, config)

	return NewTwitchFetchLiveVideosUsecase(
		liveVideoRepository,
		archiveVideoRepository,
		twitchApiClient,
		time.Now,
	)
}

func NewInjectTwitchUpdateLiveVideosUsecase(ctx context.Context) usecaseBatch.TwitchUpdateLiveVideosUsecaseInterface {
	config := cfg.NewConfig()
	postgresql := postgresql.NewPostgreSQL(config)

	archiveVideoRepository := postgresqlLive.NewArchiveVideoRepository(postgresql)

	client := httpClient.NewHttpClient(nil)
	twitchApiClient := twitch.NewTwitchApiClient(client, config)

	return NewTwitchUpdateLiveVideosUsecase(
		archiveVideoRepository,
		twitchApiClient,
		time.Now,
	)
}
