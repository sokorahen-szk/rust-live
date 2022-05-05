package live

import (
	"context"
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/domain/common"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
	"github.com/stretchr/testify/assert"
)

func Test_ArchiveVideoRepository_Create(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	postgresql := postgresql.NewPostgreSQL(cfg.NewConfig())
	repository := NewArchiveVideoRepository(postgresql)

	tm := common.NewDatetime("2022-02-02T14:00:00Z")
	in := &input.ArchiveVideoInput{
		BroadcastId:     "39300467239",
		Title:           "title",
		Url:             "https://example.com/test",
		Stremer:         "テスター",
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

	t.Run("BroadcastIdでデータが見つかる場合、データが返されること", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})

		searchBroadcastId := entity.NewVideoBroadcastId("39300467239")

		datetime := common.NewDatetime("2022-02-02T14:00:00Z")
		in := &input.ArchiveVideoInput{
			BroadcastId:     "39300467239",
			Title:           "title",
			Url:             "https://example.com/test",
			Stremer:         "テスター",
			ThumbnailImage:  "https://example.com/test.jpg",
			StartedDatetime: datetime.Time(),
		}
		err := repository.Create(ctx, in)
		a.NoError(err)

		actual, err := repository.GetByBroadcastId(ctx, searchBroadcastId)
		a.NoError(err)
		a.NotNil(actual)
		a.Equal(searchBroadcastId.String(), actual.GetBroadcastId().String())
		a.Equal("title", actual.GetTitle().String())
		a.Equal("https://example.com/test", actual.GetUrl().String())
		a.Equal("テスター", actual.GetStremer().String())
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
