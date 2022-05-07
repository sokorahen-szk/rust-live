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
	generateArchiveVideo(t, ctx, "39300467239", datetime, repository)

	list, err := repository.List(ctx)
	a.NoError(err)
}

func generateArchiveVideo(t *testing.T, ctx context.Context, broadcastId string,
	datetime *common.Datetime, repo repoIf.ArchiveVideoRepositoryInterface) *input.ArchiveVideoInput {
	platform := entity.NewPlatform(entity.PlatformTwitch)
	status := entity.NewVideoStatus(entity.VideoStatusStreaming)
	in := &input.ArchiveVideoInput{
		BroadcastId:     broadcastId,
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

	return in
}
