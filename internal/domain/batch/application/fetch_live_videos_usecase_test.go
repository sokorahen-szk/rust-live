package application

import (
	"context"
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	redisLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
	"github.com/stretchr/testify/assert"
)

func Test_FetchLiveVideosUsecase_Handle(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	postgresql := postgresql.NewPostgreSQL(cfg.NewConfig())
	redis := redis.NewRedis(ctx, cfg.NewConfig())

	liveVideoRepository := redisLive.NewLiveVideoRepository(redis)

	postgresql.Truncate([]string{"archive_videos"})
	redis.Truncate()

	listLiveUsecase := NewInjectFetchLiveVideosUsecase(ctx)
	err := listLiveUsecase.Handle(ctx)
	a.NoError(err)

	_, err = liveVideoRepository.List(ctx, &list.ListLiveVideosInput{})
	a.NoError(err)
}
