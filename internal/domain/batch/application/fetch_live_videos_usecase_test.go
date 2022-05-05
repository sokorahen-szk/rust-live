package application

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch/twitch"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
	postgresqlLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql/live"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	redisLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
	testTwitch "github.com/sokorahen-szk/rust-live/tests/infrastructure/batch/twitch"
)

func Test_FetchLiveVideosUsecase_Handle(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	config := cfg.NewConfig()
	redis := redis.NewRedis(ctx, config)
	postgresql := postgresql.NewPostgreSQL(config)

	archiveVideoRepository := postgresqlLive.NewArchiveVideoRepository(postgresql)
	liveVideoRepository := redisLive.NewLiveVideoRepository(redis)

	listBroadcastOptions := []batch.RequestParam{
		{Key: "language", Value: "ja"},
		{Key: "game_id", Value: twitch.RustGameId},
		{Key: "type", Value: "live"},
		{Key: "first", Value: "100"},
	}

	listVideoByUserIdOptions := []batch.RequestParam{
		{Key: "first", Value: "1"},
	}

	userId := "23456"

	listBroadcastResponse := &twitch.ListBroadcastResponse{
		List: []twitch.ListBroadcast{
			{
				StreamId:     "123456789",
				UserId:       "23456",
				UserLogin:    "tester",
				UserName:     "テスター",
				Title:        "テスト配信",
				ViewerCount:  12,
				StartedAt:    "2022-02-02T00:00:00Z",
				ThumbnailUrl: "https://example.com/test.jpg",
			},
		},
	}

	listVideoByUserIdResponse := &twitch.ListVideoByUserIdResponse{
		List: []twitch.ListVideoByUserId{
			{
				Id:       "234567891",
				StreamId: "123456789",
				UserId:   "23456",
				UserName: "テスター",
				Title:    "テスト配信",
				Viewable: "public",
			},
		},
	}

	t.Run("正常系/配信中の動画・ビデオ公開している場合、すべての処理が正常に完了し、エラーを返さないこと", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})
		redis.Truncate()

		ctrl := gomock.NewController(t)
		mockTwitchApiClient := testTwitch.NewMockTwitchApiClientInterface(ctrl)

		gomock.InOrder(
			mockTwitchApiClient.EXPECT().
				ListBroadcast(gomock.Eq(listBroadcastOptions)).
				Do(func(params []batch.RequestParam) {
					a.Equal(listBroadcastOptions, params)
				}).
				Return(listBroadcastResponse, nil).
				Times(1),

			mockTwitchApiClient.EXPECT().
				ListVideoByUserId(gomock.Eq(userId), gomock.Eq(listVideoByUserIdOptions)).
				Do(func(actualUserId string, params []batch.RequestParam) {
					a.Equal(listVideoByUserIdOptions, params)
					a.Equal(userId, actualUserId)
				}).
				Return(listVideoByUserIdResponse, nil).
				Times(1),
		)

		usecase := NewFetchLiveVideosUsecase(
			liveVideoRepository,
			archiveVideoRepository,
			mockTwitchApiClient,
		)

		err := usecase.Handle(ctx)
		a.NoError(err)

		broadcastId := entity.NewVideoBroadcastId("234567891")
		actualArchiveVideo, err := archiveVideoRepository.GetByBroadcastId(ctx, broadcastId)
		a.NoError(err)
		a.Equal("234567891", actualArchiveVideo.GetBroadcastId().String())
		a.Equal("テスト配信", actualArchiveVideo.GetTitle().String())
		a.Equal("https://www.twitch.tv/videos/234567891", actualArchiveVideo.GetUrl().String())
		a.Equal("テスター", actualArchiveVideo.GetStremer().String())
		a.Equal("https://example.com/test.jpg", actualArchiveVideo.GetThumbnailImage().String())
		a.Equal("2022-02-02T00:00:00Z", actualArchiveVideo.GetStartedDatetime().RFC3339())
		a.Nil(actualArchiveVideo.GetEndedDatetime())
	})
}
