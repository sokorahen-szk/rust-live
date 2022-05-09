package live

import (
	"context"
	"database/sql"
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/domain/common"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
	repoIf "github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
	"github.com/stretchr/testify/assert"
)

func Test_ArchiveVideoRepository_Create(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	postgresql := postgresql.NewPostgreSQL(cfg.NewConfig())
	repository := NewArchiveVideoRepository(postgresql)

	platform := entity.NewPlatform(entity.PlatformTwitch)

	tm := common.NewDatetime("2022-02-02T14:00:00Z")
	in := &input.ArchiveVideoInput{
		BroadcastId:     "39300467239",
		Title:           "title",
		Url:             &sql.NullString{String: "https://example.com/test", Valid: true},
		Stremer:         "テスター",
		Platform:        platform.Int(),
		ThumbnailImage:  "https://example.com/test.jpg",
		StartedDatetime: tm.Time(),
	}
	t.Run("正常に1件追加できること", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})

		err := repository.Create(ctx, in)
		a.NoError(err)
	})
	t.Run("BroadcastIdが既に登録されているとき、エラーで書き込みできないこと", func(t *testing.T) {
		err := repository.Create(ctx, in)
		a.Error(err)
	})
}

func Test_ArchiveVideoRepository_GetByBroadcastId(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	postgresql := postgresql.NewPostgreSQL(cfg.NewConfig())
	repository := NewArchiveVideoRepository(postgresql)

	platform := entity.NewPlatform(entity.PlatformTwitch)
	status := entity.NewVideoStatus(entity.VideoStatusStreaming)

	t.Run("BroadcastIdでデータが見つかる場合、データが返されること", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})

		searchBroadcastId := entity.NewVideoBroadcastId("39300467239")

		datetime := common.NewDatetime("2022-02-02T14:00:00Z")
		in := &input.ArchiveVideoInput{
			BroadcastId:     "39300467239",
			Title:           "title",
			Url:             &sql.NullString{String: "https://example.com/test", Valid: true},
			Stremer:         "テスター",
			Platform:        platform.Int(),
			Status:          status.Int(),
			ThumbnailImage:  "https://example.com/test.jpg",
			StartedDatetime: datetime.Time(),
		}
		err := repository.Create(ctx, in)
		a.NoError(err)
		// GORMの機能により、構造体のポインタを渡しているため、repository.Create
		// の中でIdがセットされて返される。
		a.NotNil(in.Id)

		actual, err := repository.GetByBroadcastId(ctx, searchBroadcastId)
		a.NoError(err)
		a.NotNil(actual)
		a.Equal(searchBroadcastId.String(), actual.GetBroadcastId().String())
		a.Equal("title", actual.GetTitle().String())
		a.Equal("https://example.com/test", actual.GetUrl().String())
		a.Equal("テスター", actual.GetStremer().String())
		a.Equal(1, platform.Int())
		a.Equal("https://example.com/test.jpg", actual.GetThumbnailImage().String())
		a.Equal(datetime.RFC3339(), actual.GetStartedDatetime().RFC3339())
		a.Nil(actual.GetEndedDatetime())
	})
	t.Run("BroadcastIdでデータが見つからない場合、エラーが返されること", func(t *testing.T) {
		searchBroadcastId := entity.NewVideoBroadcastId("99900227123")
		actual, err := repository.GetByBroadcastId(ctx, searchBroadcastId)
		a.Error(err)
		a.Nil(actual)
	})
}
func Test_ArchiveVideoRepository_List(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	postgresql := postgresql.NewPostgreSQL(cfg.NewConfig())
	repository := NewArchiveVideoRepository(postgresql)

	datetime := common.NewDatetime("2022-02-02T14:00:00Z")

	t.Run("listオプションなしの場合、すべて取得できること", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})

		generateArchiveVideo(t, ctx, "39300467239", datetime, entity.VideoStatusStreaming, repository)
		generateArchiveVideo(t, ctx, "39300467240", datetime, entity.VideoStatusStreaming, repository)
		generateArchiveVideo(t, ctx, "39300467241", datetime, entity.VideoStatusEnded, repository)

		listInput := &input.ListArchiveVideoInput{}

		list, err := repository.List(ctx, listInput)
		a.NoError(err)
		a.Len(list, 3)
	})
	t.Run("listオプション、ステータス(Streaming = 1)で配信中の動画のみ取得できること", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})

		generateArchiveVideo(t, ctx, "39300467239", datetime, entity.VideoStatusStreaming, repository)
		generateArchiveVideo(t, ctx, "39300467240", datetime, entity.VideoStatusStreaming, repository)
		generateArchiveVideo(t, ctx, "39300467241", datetime, entity.VideoStatusEnded, repository)

		searchVideoStatusStreaming := entity.NewVideoStatus(entity.VideoStatusStreaming)

		listInput := &input.ListArchiveVideoInput{
			VideoStatuses: []int{searchVideoStatusStreaming.Int()},
		}

		list, err := repository.List(ctx, listInput)
		a.NoError(err)
		a.Len(list, 2)
	})
	t.Run("取得できるデータがないとき、空を返すこと", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})

		listInput := &input.ListArchiveVideoInput{}

		list, err := repository.List(ctx, listInput)
		a.NoError(err)
		a.Len(list, 0)
	})
}

func Test_ArchiveVideoRepository_Update(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	postgresql := postgresql.NewPostgreSQL(cfg.NewConfig())
	repository := NewArchiveVideoRepository(postgresql)

	datetime := common.NewDatetime("2022-02-02T14:00:00Z")

	t.Run("ステータス(Streaming = 1)をステータス(Ended=2)に変更できること", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})

		expectedArchiveVideo := generateArchiveVideo(t, ctx, "39300467239", datetime, entity.VideoStatusStreaming, repository)
		updateVideoStatus := entity.NewVideoStatus(entity.VideoStatusEnded)
		updateInput := &input.UpdateArchiveVideoInput{
			Status: updateVideoStatus,
		}
		err := repository.Update(ctx, expectedArchiveVideo.GetId(), updateInput)
		a.NoError(err)

		broadcastId := entity.NewVideoBroadcastId("39300467239")
		actualArchiveVideo, err := repository.GetByBroadcastId(ctx, broadcastId)
		a.NoError(err)
		a.Equal(expectedArchiveVideo.GetStatus(), actualArchiveVideo.GetStatus())

	})
	t.Run("ステータス(Streaming = 1)をステータス(Streaming=1)に変更してもエラーにならないこと", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})

		expectedArchiveVideo := generateArchiveVideo(t, ctx, "39300467239", datetime, entity.VideoStatusStreaming, repository)
		updateVideoStatus := entity.NewVideoStatus(entity.VideoStatusStreaming)
		updateInput := &input.UpdateArchiveVideoInput{
			Status: updateVideoStatus,
		}
		err := repository.Update(ctx, expectedArchiveVideo.GetId(), updateInput)
		a.NoError(err)

		broadcastId := entity.NewVideoBroadcastId("39300467239")
		actualArchiveVideo, err := repository.GetByBroadcastId(ctx, broadcastId)
		a.NoError(err)
		a.Equal(expectedArchiveVideo.GetStatus(), actualArchiveVideo.GetStatus())
	})
	t.Run("更新するデータが見つからない場合、エラーになること", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})

		expectedArchiveVideo := generateArchiveVideo(t, ctx, "39300467239", datetime, entity.VideoStatusStreaming, repository)
		updateVideoStatus := entity.NewVideoStatus(entity.VideoStatusEnded)
		updateInput := &input.UpdateArchiveVideoInput{
			Status: updateVideoStatus,
		}

		notFoundArchiveVideoId := entity.NewVideoId(98765)

		err := repository.Update(ctx, notFoundArchiveVideoId, updateInput)
		a.Error(err)

		broadcastId := entity.NewVideoBroadcastId("39300467239")
		actualArchiveVideo, err := repository.GetByBroadcastId(ctx, broadcastId)
		a.NoError(err)
		a.Equal(expectedArchiveVideo.GetStatus(), actualArchiveVideo.GetStatus())
	})
}

func generateArchiveVideo(t *testing.T, ctx context.Context, broadcastId string,
	datetime *common.Datetime, videoStatus entity.VideoStatus, repo repoIf.ArchiveVideoRepositoryInterface) *entity.ArchiveVideo {
	platform := entity.NewPlatform(entity.PlatformTwitch)
	status := entity.NewVideoStatus(videoStatus)
	videoBroadcastId := entity.NewVideoBroadcastId(broadcastId)
	in := &input.ArchiveVideoInput{
		BroadcastId:     videoBroadcastId.String(),
		Title:           "title",
		Url:             &sql.NullString{String: "https://example.com/test", Valid: true},
		Stremer:         "テスター",
		Platform:        platform.Int(),
		Status:          status.Int(),
		ThumbnailImage:  "https://example.com/test.jpg",
		StartedDatetime: datetime.Time(),
	}

	err := repo.Create(ctx, in)
	assert.NoError(t, err)
	// GORMの機能により、構造体のポインタを渡しているため、repository.Create
	// の中でIdがセットされて返される。
	assert.NotNil(t, in.Id)

	archiveVideo, err := repo.GetByBroadcastId(ctx, videoBroadcastId)
	assert.NoError(t, err)

	return archiveVideo
}
