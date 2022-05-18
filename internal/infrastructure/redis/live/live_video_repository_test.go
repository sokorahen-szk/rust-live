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
	var (
		correctSort  *entity.LiveVideoSortKey = entity.NewLiveVideoSortKey(0)
		correctPage  int                      = 1
		correctLimit int                      = 10
	)

	a := assert.New(t)
	ctx := context.Background()
	c := cfg.NewConfig()
	redis := redis.NewRedis(ctx, c)
	liveVideoRepository := NewLiveVideoRepository(redis)

	t.Run("SearchKeywordsで検索でき、データが取得できること", func(t *testing.T) {
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
			input := list.NewListLiveVideoInput(
				p.arg,
				correctSort,
				correctPage,
				correctLimit,
			)

			t.Run(p.name, func(t *testing.T) {
				actual, err := liveVideoRepository.List(ctx, input)
				a.Len(actual, p.want)
				a.NoError(err)
			})
		}
	})
	t.Run("Sortでデータの順番が変わり、意図したデータの並びで取得できること", func(t *testing.T) {
		redis.Truncate()

		liveVideo := mockEntity.NewMockLiveVideo(1)
		liveVideo.Viewer = entity.NewVideoViewer(99)
		liveVideo2 := mockEntity.NewMockLiveVideo(2)
		liveVideo2.Viewer = entity.NewVideoViewer(101)
		liveVideo3 := mockEntity.NewMockLiveVideo(3)
		liveVideo3.Viewer = entity.NewVideoViewer(100)

		liveVideos := []*entity.LiveVideo{
			liveVideo,
			liveVideo2,
			liveVideo3,
		}

		err := liveVideoRepository.Create(ctx, liveVideos)
		a.NoError(err)

		t.Run("viewerで昇順検索が正しくできること", func(t *testing.T) {
			listInput := list.NewListLiveVideoInput("", entity.NewLiveVideoSortKey(1), 1, 0)
			actualListVideos, err := liveVideoRepository.List(ctx, listInput)
			a.NoError(err)
			a.Equal(liveVideo.GetViewer(), actualListVideos[0].GetViewer())
			a.Equal(liveVideo3.GetViewer(), actualListVideos[2].GetViewer())
			a.Equal(liveVideo2.GetViewer(), actualListVideos[1].GetViewer())
		})
		t.Run("viewerで降順検索が正しくできること", func(t *testing.T) {
			listInput := list.NewListLiveVideoInput("", entity.NewLiveVideoSortKey(2), 1, 0)
			actualListVideos, err := liveVideoRepository.List(ctx, listInput)
			a.NoError(err)
			a.Equal(liveVideo2.GetViewer(), actualListVideos[1].GetViewer())
			a.Equal(liveVideo3.GetViewer(), actualListVideos[2].GetViewer())
			a.Equal(liveVideo.GetViewer(), actualListVideos[0].GetViewer())
		})
	})
	t.Run("Pageでページ送りができ、意図したページのデータが取得できること", func(t *testing.T) {
		// TODO:
	})
	t.Run("Limitで指定した件数でデータが取得できること", func(t *testing.T) {
		// TODO:
	})
	t.Run("redis内にデータが1件も存在しない場合、空配列を返すこと", func(t *testing.T) {
		redis.Truncate()

		liveVideos, err := liveVideoRepository.List(ctx, &list.ListLiveVideoInput{})
		a.NoError(err)
		a.Len(liveVideos, 0)
	})
}
