package application

import (
	"context"
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	redisLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
	"github.com/stretchr/testify/assert"
)

func Test_ListLiveVideosUsecase_Handle(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	searchKeywords := "keywords"
	platforms := []*entity.Platform{}
	sortKey := entity.NewLiveVideoSortKey(0)

	config := cfg.NewConfig()
	redis := redis.NewRedis(ctx, config)
	liveVideoRepository := redisLive.NewLiveVideoRepository(redis)

	usecase := NewListLiveVideosUsecase(liveVideoRepository)

	t.Run("liveVideoが0件の時、空を返すこと", func(t *testing.T) {
		input := list.NewListLiveVideoInput(searchKeywords, platforms, sortKey, 1, 10)

		actual, err := usecase.Handle(ctx, input)
		a.NotNil(actual)
		a.NoError(err)
		a.Len(actual.LiveVideos, 0)
	})
}
