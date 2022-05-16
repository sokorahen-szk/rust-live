package application

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/domain/common"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch/twitch"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
	postgresqlLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql/live"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	redisLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
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

	now := common.NewDatetime("2022-02-02T00:01:32Z")
	sortKey := entity.NewLiveVideoSortKey(0)

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
	expectedElapsedTimes := 92

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
			now.TimeFunc(),
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

		listInput := list.NewListLiveVideoInput("", sortKey, 1, 0)
		actualLiveVideos, err := liveVideoRepository.List(ctx, listInput)
		a.NoError(err)
		a.Len(actualLiveVideos, 1)
		a.Equal(actualArchiveVideo.GetId().Int(), actualLiveVideos[0].GetId().Int())
		a.Equal("234567891", actualLiveVideos[0].GetBroadcastId().String())
		a.Equal("テスト配信", actualLiveVideos[0].GetTitle().String())
		a.Equal("https://www.twitch.tv/tester", actualLiveVideos[0].GetUrl().String())
		a.Equal("テスター", actualLiveVideos[0].GetStremer().String())
		a.Equal(12, actualLiveVideos[0].GetViewer().Int())
		a.Equal("https://example.com/test.jpg", actualLiveVideos[0].GetThumbnailImage().String())
		a.Equal("2022-02-02T00:00:00Z", actualLiveVideos[0].GetStartedDatetime().RFC3339())
		a.Equal(expectedElapsedTimes, actualLiveVideos[0].GetElapsedTimes().Int())
	})
	t.Run("正常系/配信中の動画・ビデオ公開がない場合、すべての処理が正常に完了し、エラーを返さないこと", func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})
		redis.Truncate()

		ctrl := gomock.NewController(t)
		mockTwitchApiClient := testTwitch.NewMockTwitchApiClientInterface(ctrl)

		emptyListVideoByUserIdResponse := &twitch.ListVideoByUserIdResponse{
			List: []twitch.ListVideoByUserId{},
		}

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
				Return(emptyListVideoByUserIdResponse, nil).
				Times(1),
		)

		usecase := NewFetchLiveVideosUsecase(
			liveVideoRepository,
			archiveVideoRepository,
			mockTwitchApiClient,
			now.TimeFunc(),
		)

		err := usecase.Handle(ctx)
		a.NoError(err)

		broadcastId := entity.NewVideoBroadcastId("123456789")
		actualArchiveVideo, err := archiveVideoRepository.GetByBroadcastId(ctx, broadcastId)
		a.NoError(err)
		a.Equal("123456789", actualArchiveVideo.GetBroadcastId().String())
		a.Equal("テスト配信", actualArchiveVideo.GetTitle().String())
		a.Nil(actualArchiveVideo.GetUrl())
		a.Equal("テスター", actualArchiveVideo.GetStremer().String())
		a.Equal("https://example.com/test.jpg", actualArchiveVideo.GetThumbnailImage().String())
		a.Equal("2022-02-02T00:00:00Z", actualArchiveVideo.GetStartedDatetime().RFC3339())
		a.Nil(actualArchiveVideo.GetEndedDatetime())
	})
	t.Run("異常系/twitchAPI ListBroadcastの取得に失敗した場合、エラーで終了すること", func(t *testing.T) {
		// TODO:
	})
	t.Run("異常系/twitchAPI ListVideoByUserIdの取得に失敗した場合、エラーで終了すること", func(t *testing.T) {
		// TODO:
	})
}
