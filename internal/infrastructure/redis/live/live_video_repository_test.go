package live

import (
	"context"
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
	"github.com/stretchr/testify/assert"

	mockEntity "github.com/sokorahen-szk/rust-live/tests/domain/live/entity"
)

func Test_LiveVideoRepository_Create(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	c := cfg.NewConfig()
	redis := redis.NewRedis(ctx, c)
	liveVideoRepository := NewLiveVideoRepository(redis)

	t.Run("LiveVideo型の複数データをRedisに書き込めること", func(t *testing.T) {
		liveVideos := []*entity.LiveVideo{
			mockEntity.NewMockLiveVideo(1),
			mockEntity.NewMockLiveVideo(2),
			mockEntity.NewMockLiveVideo(3),
		}
		err := liveVideoRepository.Create(ctx, liveVideos)
		a.NoError(err)

		actualLiveVideos, err := liveVideoRepository.List(ctx, &list.ListLiveVideosInput{})
		a.NoError(err)
		a.Len(actualLiveVideos, 3)
	})

	t.Run("LiveVideo型のデータが空配列指定の時、空配列を書き込むこと", func(t *testing.T) {
		emptyLiveVideos := []*entity.LiveVideo{}
		err := liveVideoRepository.Create(ctx, emptyLiveVideos)
		a.NoError(err)

		actualLiveVideos, err := liveVideoRepository.List(ctx, &list.ListLiveVideosInput{})
		a.NoError(err)
		a.Len(actualLiveVideos, 0)
	})
}

func Test_LiveVideoRepository_List(t *testing.T) {
	//
}
