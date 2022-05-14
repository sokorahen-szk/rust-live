package application

import (
	"context"
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/domain/common"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
	postgresqlLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql/live"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	redisLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
	"github.com/stretchr/testify/assert"

	mockEntity "github.com/sokorahen-szk/rust-live/tests/domain/live/entity"
	mockInput "github.com/sokorahen-szk/rust-live/tests/domain/live/input"
)

func Test_UpdateLiveVideosUsecase_Handle(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	config := cfg.NewConfig()
	redis := redis.NewRedis(ctx, config)
	postgresql := postgresql.NewPostgreSQL(config)

	archiveVideoRepository := postgresqlLive.NewArchiveVideoRepository(postgresql)
	liveVideoRepository := redisLive.NewLiveVideoRepository(redis)

	platformTwitch := entity.NewPlatform(entity.PlatformTwitch)
	videoStatusStreaming := entity.NewVideoStatus(entity.VideoStatusStreaming)

	now := common.NewDatetime("2022-01-01T15:00:00Z")
	startedDatetime := entity.NewStartedDatetimeFromTime(now.Time())

	url := "https://example.com/test"

	t.Run(`正常系/archive_videosに動画情報があり、
		live_videosにない場合、archive_videosのstatus = 2に変更する`, func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})
		redis.Truncate()

		in := mockInput.NewMockArchiveVideoInput(1, url, platformTwitch, videoStatusStreaming, startedDatetime, nil)
		err := archiveVideoRepository.Create(ctx, in)
		a.NoError(err)

		usecase := NewUpdateLiveVideosUsecase(
			liveVideoRepository,
			archiveVideoRepository,
			now.TimeFunc(),
		)

		err = usecase.Handle(ctx)
		a.NoError(err)
	})
	t.Run("正常系/archive_videos・live_videosに動画情報がある場合、statusを変更しない", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})
		redis.Truncate()

		mockLiveVideoIn := mockInput.NewMockArchiveVideoInput(1, url, platformTwitch, videoStatusStreaming, startedDatetime, nil)
		err := archiveVideoRepository.Create(ctx, mockLiveVideoIn)
		a.NoError(err)

		liveVideo := mockEntity.NewMockLiveVideo(mockLiveVideoIn.Id)
		err = liveVideoRepository.Create(ctx, []*entity.LiveVideo{liveVideo})
		a.NoError(err)
	})
}
