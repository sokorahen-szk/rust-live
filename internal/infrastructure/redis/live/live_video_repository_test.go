package live

import (
	"context"
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
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
		err := liveVideoRepository.Create(ctx, liveVideos, repository.TwitchLiveVideoKey)
		a.NoError(err)

		actualLiveVideos, err := liveVideoRepository.List(ctx, &list.ListLiveVideoInput{}, repository.TwitchLiveVideoKey)
		a.NoError(err)
		a.Len(actualLiveVideos, 3)
	})

	t.Run("LiveVideo型のデータが空配列指定の時、空配列を書き込むこと", func(t *testing.T) {
		redis.Truncate()

		emptyLiveVideos := []*entity.LiveVideo{}
		err := liveVideoRepository.Create(ctx, emptyLiveVideos, repository.TwitchLiveVideoKey)
		a.NoError(err)

		actualLiveVideos, err := liveVideoRepository.List(ctx, &list.ListLiveVideoInput{}, repository.TwitchLiveVideoKey)
		a.NoError(err)
		a.Len(actualLiveVideos, 0)
	})
}

func Test_LiveVideoRepository_List(t *testing.T) {
	var (
		correctSort  *entity.LiveVideoSortKey = entity.NewLiveVideoSortKey(0)
		correctPage  int                      = 1
		correctLimit int                      = 10

		correctPlatforms []*entity.Platform = []*entity.Platform{
			entity.NewPlatform(entity.PlatformTwitch),
		}
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

		err := liveVideoRepository.Create(ctx, liveVideos, repository.TwitchLiveVideoKey)
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
				correctPlatforms,
				correctSort,
				correctPage,
				correctLimit,
			)

			t.Run(p.name, func(t *testing.T) {
				actual, err := liveVideoRepository.List(ctx, input, repository.TwitchLiveVideoKey)
				a.Len(actual, p.want)
				a.NoError(err)
			})
		}
	})
	t.Run("Sortでデータの順番が変わり、意図したデータの並びで取得できること", func(t *testing.T) {
		redis.Truncate()

		liveVideo := mockEntity.NewMockLiveVideo(1)
		liveVideo.Viewer = entity.NewVideoViewer(99)
		liveVideo.StartedDatetime = entity.NewStartedDatetime("2022-01-01T00:00:01Z")
		liveVideo2 := mockEntity.NewMockLiveVideo(2)
		liveVideo2.Viewer = entity.NewVideoViewer(101)
		liveVideo2.StartedDatetime = entity.NewStartedDatetime("2021-12-31T23:59:59Z")
		liveVideo3 := mockEntity.NewMockLiveVideo(3)
		liveVideo3.Viewer = entity.NewVideoViewer(100)
		liveVideo3.StartedDatetime = entity.NewStartedDatetime("2022-01-01T00:00:00Z")

		liveVideos := []*entity.LiveVideo{
			liveVideo,
			liveVideo2,
			liveVideo3,
		}

		err := liveVideoRepository.Create(ctx, liveVideos, repository.TwitchLiveVideoKey)
		a.NoError(err)

		t.Run("viewerで昇順検索が正しくできること", func(t *testing.T) {
			listInput := list.NewListLiveVideoInput("", correctPlatforms, entity.NewLiveVideoSortKey(entity.LiveVideoViewerAsc), 1, 0)
			actualListVideos, err := liveVideoRepository.List(ctx, listInput, repository.TwitchLiveVideoKey)
			a.NoError(err)
			a.Equal(liveVideo.GetViewer(), actualListVideos[0].GetViewer())
			a.Equal(liveVideo3.GetViewer(), actualListVideos[1].GetViewer())
			a.Equal(liveVideo2.GetViewer(), actualListVideos[2].GetViewer())
		})
		t.Run("viewerで降順検索が正しくできること", func(t *testing.T) {
			listInput := list.NewListLiveVideoInput("", correctPlatforms, entity.NewLiveVideoSortKey(entity.LiveVideoViewerDesc), 1, 0)
			actualListVideos, err := liveVideoRepository.List(ctx, listInput, repository.TwitchLiveVideoKey)
			a.NoError(err)
			a.Equal(liveVideo2.GetViewer(), actualListVideos[0].GetViewer())
			a.Equal(liveVideo3.GetViewer(), actualListVideos[1].GetViewer())
			a.Equal(liveVideo.GetViewer(), actualListVideos[2].GetViewer())
		})
		t.Run("startedDatetimeで昇順検索が正しくできること", func(t *testing.T) {
			listInput := list.NewListLiveVideoInput("", correctPlatforms, entity.NewLiveVideoSortKey(entity.LiveVideoStartedDatetimeAsc), 1, 0)
			actualListVideos, err := liveVideoRepository.List(ctx, listInput, repository.TwitchLiveVideoKey)
			a.NoError(err)
			a.Equal(liveVideo2.GetStartedDatetime(), actualListVideos[0].GetStartedDatetime())
			a.Equal(liveVideo3.GetStartedDatetime(), actualListVideos[1].GetStartedDatetime())
			a.Equal(liveVideo.GetStartedDatetime(), actualListVideos[2].GetStartedDatetime())
		})
		t.Run("startedDatetimeで降順検索が正しくできること", func(t *testing.T) {
			listInput := list.NewListLiveVideoInput("", correctPlatforms, entity.NewLiveVideoSortKey(entity.LiveVideoStartedDatetimeDesc), 1, 0)
			actualListVideos, err := liveVideoRepository.List(ctx, listInput, repository.TwitchLiveVideoKey)
			a.NoError(err)
			a.Equal(liveVideo.GetStartedDatetime(), actualListVideos[0].GetStartedDatetime())
			a.Equal(liveVideo3.GetStartedDatetime(), actualListVideos[1].GetStartedDatetime())
			a.Equal(liveVideo2.GetStartedDatetime(), actualListVideos[2].GetStartedDatetime())
		})
	})
	t.Run("Pageでページ送りができ、意図したページのデータが取得できること", func(t *testing.T) {
		redis.Truncate()

		liveVideo := mockEntity.NewMockLiveVideo(1)
		liveVideo2 := mockEntity.NewMockLiveVideo(2)
		liveVideo3 := mockEntity.NewMockLiveVideo(3)

		liveVideos := []*entity.LiveVideo{
			liveVideo,
			liveVideo2,
			liveVideo3,
		}

		err := liveVideoRepository.Create(ctx, liveVideos, repository.TwitchLiveVideoKey)
		a.NoError(err)

		t.Run("3件中2件毎に取得した場合、2ページ分のデータが取得できること", func(t *testing.T) {
			listInput := list.NewListLiveVideoInput("", correctPlatforms, entity.NewLiveVideoSortKey(1), 1, 2)
			actualListVideos, err := liveVideoRepository.List(ctx, listInput, repository.TwitchLiveVideoKey)
			a.NoError(err)
			a.Len(actualListVideos, 2)
			a.Equal(liveVideo.GetId(), actualListVideos[0].GetId())
			a.Equal(liveVideo2.GetId(), actualListVideos[1].GetId())

			listInput = list.NewListLiveVideoInput("", correctPlatforms, entity.NewLiveVideoSortKey(1), 2, 2)
			actualListVideos, err = liveVideoRepository.List(ctx, listInput, repository.TwitchLiveVideoKey)
			a.NoError(err)
			a.Len(actualListVideos, 1)
			a.Equal(liveVideo3.GetId(), actualListVideos[0].GetId())
		})
		t.Run("3件中3件毎に取得した場合、1ページ分のデータが取得できること", func(t *testing.T) {
			listInput := list.NewListLiveVideoInput("", correctPlatforms, entity.NewLiveVideoSortKey(1), 1, 3)
			actualListVideos, err := liveVideoRepository.List(ctx, listInput, repository.TwitchLiveVideoKey)
			a.NoError(err)
			a.Len(actualListVideos, 3)
			a.Equal(liveVideo.GetId(), actualListVideos[0].GetId())
			a.Equal(liveVideo2.GetId(), actualListVideos[1].GetId())
			a.Equal(liveVideo3.GetId(), actualListVideos[2].GetId())

			listInput = list.NewListLiveVideoInput("", correctPlatforms, entity.NewLiveVideoSortKey(1), 2, 3)
			actualListVideos, err = liveVideoRepository.List(ctx, listInput, repository.TwitchLiveVideoKey)
			a.NoError(err)
			a.Len(actualListVideos, 0)
		})
	})
	t.Run("Limitで指定した件数でデータが取得できること", func(t *testing.T) {
		redis.Truncate()

		liveVideo := mockEntity.NewMockLiveVideo(1)
		liveVideo2 := mockEntity.NewMockLiveVideo(2)
		liveVideo3 := mockEntity.NewMockLiveVideo(3)

		liveVideos := []*entity.LiveVideo{
			liveVideo,
			liveVideo2,
			liveVideo3,
		}

		err := liveVideoRepository.Create(ctx, liveVideos, repository.TwitchLiveVideoKey)
		a.NoError(err)

		listInput := list.NewListLiveVideoInput("", correctPlatforms, entity.NewLiveVideoSortKey(1), 1, 3)
		actualListVideos, err := liveVideoRepository.List(ctx, listInput, repository.TwitchLiveVideoKey)
		a.NoError(err)
		a.Len(actualListVideos, 3)
		a.Equal(liveVideo.GetId(), actualListVideos[0].GetId())
		a.Equal(liveVideo2.GetId(), actualListVideos[1].GetId())
		a.Equal(liveVideo3.GetId(), actualListVideos[2].GetId())
	})

	t.Run("Platformsで検索するプラットフォームを指定し、意図したデータが返されること", func(t *testing.T) {
		redis.Truncate()

		liveVideo := mockEntity.NewMockLiveVideo(1)
		liveVideo2 := mockEntity.NewMockLiveVideo(2)
		liveVideo3 := mockEntity.NewMockLiveVideo(3)
		liveVideo3.Platform = entity.NewPlatform(entity.PlatformYoutube)

		liveVideos := []*entity.LiveVideo{
			liveVideo,
			liveVideo2,
			liveVideo3,
		}

		err := liveVideoRepository.Create(ctx, liveVideos, repository.TwitchLiveVideoKey)
		a.NoError(err)

		t.Run("プラットフォームをtwitchに指定し、2件取得できること", func(t *testing.T) {
			listInput := list.NewListLiveVideoInput("", correctPlatforms, entity.NewLiveVideoSortKey(1), 1, 10)
			actualListVideos, err := liveVideoRepository.List(ctx, listInput, repository.TwitchLiveVideoKey)
			a.NoError(err)
			a.Len(actualListVideos, 2)
			a.Equal(liveVideo.GetId(), actualListVideos[0].GetId())
			a.Equal(liveVideo2.GetId(), actualListVideos[1].GetId())
		})
		t.Run("プラットフォームをyoutubeに指定し、1件取得できること", func(t *testing.T) {
			youtubePlatform := []*entity.Platform{
				entity.NewPlatform(entity.PlatformYoutube),
			}
			listInput := list.NewListLiveVideoInput("", youtubePlatform, entity.NewLiveVideoSortKey(1), 1, 10)
			actualListVideos, err := liveVideoRepository.List(ctx, listInput, repository.TwitchLiveVideoKey)
			a.NoError(err)
			a.Len(actualListVideos, 1)
			a.Equal(liveVideo3.GetId(), actualListVideos[0].GetId())
		})
	})

	t.Run("redis内にデータが1件も存在しない場合、空配列を返すこと", func(t *testing.T) {
		redis.Truncate()

		liveVideos, err := liveVideoRepository.List(ctx, &list.ListLiveVideoInput{}, repository.TwitchLiveVideoKey)
		a.NoError(err)
		a.Len(liveVideos, 0)
	})
}

func Test_LiveVideoRepository_Count(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	c := cfg.NewConfig()
	redis := redis.NewRedis(ctx, c)
	liveVideoRepository := NewLiveVideoRepository(redis)

	t.Run("0件の場合", func(t *testing.T) {
		redis.Truncate()

		count, err := liveVideoRepository.Count(ctx, repository.TwitchLiveVideoKey)
		a.NoError(err)
		a.Equal(0, count)
	})
	t.Run("5件以上の場合", func(t *testing.T) {
		redis.Truncate()

		liveVideos := []*entity.LiveVideo{
			mockEntity.NewMockLiveVideo(1),
			mockEntity.NewMockLiveVideo(2),
			mockEntity.NewMockLiveVideo(3),
			mockEntity.NewMockLiveVideo(4),
			mockEntity.NewMockLiveVideo(5),
		}

		err := liveVideoRepository.Create(ctx, liveVideos, repository.TwitchLiveVideoKey)
		a.NoError(err)

		count, err := liveVideoRepository.Count(ctx, repository.TwitchLiveVideoKey)
		a.NoError(err)
		a.Equal(5, count)
	})
}

func Test_LiveVideoRepository_Analytics(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	c := cfg.NewConfig()
	redis := redis.NewRedis(ctx, c)
	liveVideoRepository := NewLiveVideoRepository(redis)

	t.Run("0件の場合", func(t *testing.T) {
		redis.Truncate()

		output, err := liveVideoRepository.Analytics(ctx, repository.TwitchLiveVideoKey)
		a.NoError(err)
		a.Equal(0, output.GetViewers())
	})
	t.Run("5件以上の場合", func(t *testing.T) {
		redis.Truncate()

		liveVideos := []*entity.LiveVideo{
			mockEntity.NewMockLiveVideo(1),
			mockEntity.NewMockLiveVideo(2),
			mockEntity.NewMockLiveVideo(3),
			mockEntity.NewMockLiveVideo(4),
			mockEntity.NewMockLiveVideo(5),
		}

		err := liveVideoRepository.Create(ctx, liveVideos, repository.TwitchLiveVideoKey)
		a.NoError(err)

		output, err := liveVideoRepository.Analytics(ctx, repository.TwitchLiveVideoKey)
		a.NoError(err)
		a.Equal(50, output.GetViewers())
	})
}
