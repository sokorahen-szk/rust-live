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
		redis.Truncate()

		liveVideos := []*entity.LiveVideo{
			mockEntity.NewMockLiveVideo(1),
			mockEntity.NewMockLiveVideo(2),
			mockEntity.NewMockLiveVideo(3),
		}
		err := liveVideoRepository.Create(ctx, liveVideos)
		a.NoError(err)

		actualLiveVideos, err := liveVideoRepository.List(ctx, &list.ListLiveVideoInput{})
		a.NoError(err)
		a.Len(actualLiveVideos, 3)
	})

	t.Run("LiveVideo型のデータが空配列指定の時、空配列を書き込むこと", func(t *testing.T) {
		redis.Truncate()

		emptyLiveVideos := []*entity.LiveVideo{}
		err := liveVideoRepository.Create(ctx, emptyLiveVideos)
		a.NoError(err)

		actualLiveVideos, err := liveVideoRepository.List(ctx, &list.ListLiveVideoInput{})
		a.NoError(err)
		a.Len(actualLiveVideos, 0)
	})
}

func Test_LiveVideoRepository_List(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	c := cfg.NewConfig()
	redis := redis.NewRedis(ctx, c)
	liveVideoRepository := NewLiveVideoRepository(redis)
	t.Run("SearchKeywords", func(t *testing.T) {
		redis.Truncate()

		liveVideos := []*entity.LiveVideo{
			mockEntity.NewMockLiveVideo(1),
			mockEntity.NewMockLiveVideo(2),
			mockEntity.NewMockLiveVideo(3),
		}

		err := liveVideoRepository.Create(ctx, liveVideos)
		a.NoError(err)

		tests := []struct {
			name string
			arg  string
			want int
		}{
			{"検索キーワードが空の場合、3件を返す.", "", 3},
			{"検索キーワードを指定して、一致するデータがない場合、0件を返す.", "一郎", 0},
			{"検索キーワードを指定して、一致するデータが1件の場合、1件を返す.", "太郎1", 1},
		}
		for _, p := range tests {
			input := list.NewListLiveVideoInput(p.arg)

			t.Run(p.name, func(t *testing.T) {
				actual, err := liveVideoRepository.List(ctx, input)
				a.Len(actual, p.want)
				a.NoError(err)
			})
		}
	})
}
