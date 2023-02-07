package application_batch

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/domain/common"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch/twitch"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
	postgresqlLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql/live"
	httpClient "github.com/sokorahen-szk/rust-live/pkg/http"
	"github.com/stretchr/testify/assert"

	mockInput "github.com/sokorahen-szk/rust-live/tests/domain/live/input"
	testTwitch "github.com/sokorahen-szk/rust-live/tests/infrastructure/batch/twitch"
)

func Test_TwitchUpdateLiveVideosUsecase_Handle(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	config := cfg.NewConfig()
	postgresql := postgresql.NewPostgreSQL(config)

	archiveVideoRepository := postgresqlLive.NewArchiveVideoRepository(postgresql)

	platformTwitch := entity.NewPlatform(entity.PlatformTwitch)
	videoStatusStreaming := entity.NewVideoStatus(entity.VideoStatusStreaming)

	now := common.NewDatetime("2022-01-01T15:00:00Z")
	startedDatetime := entity.NewStartedDatetimeFromTime(now.Time())

	url := "https://example.com/test"

	listBroadcastOptions := []httpClient.RequestParam{
		{Key: "language", Value: "ja"},
		{Key: "game_id", Value: twitch.RustGameId},
		{Key: "type", Value: "live"},
		{Key: "first", Value: "100"},
	}

	t.Run(`正常系/archive_videosに動画情報が2件あり、
		twitchAPIのbroadcastに存在しないId = 2は、archive_videosとarchive_videosのstatus = 2に変更する`, func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})

		mockLiveVideo1In := mockInput.NewMockArchiveVideoInput(1, url, platformTwitch, videoStatusStreaming, startedDatetime, nil)
		err := archiveVideoRepository.Create(ctx, mockLiveVideo1In)
		a.NoError(err)

		mockLiveVideo2In := mockInput.NewMockArchiveVideoInput(2, url, platformTwitch, videoStatusStreaming, startedDatetime, nil)
		err = archiveVideoRepository.Create(ctx, mockLiveVideo2In)
		a.NoError(err)

		ctrl := gomock.NewController(t)
		mockTwitchApiClient := testTwitch.NewMockTwitchApiClientInterface(ctrl)

		listBroadcastResponse := &twitch.ListBroadcastResponse{
			List: []twitch.ListBroadcast{
				{
					StreamId:     mockLiveVideo1In.BroadcastId,
					UserId:       "12345",
					UserLogin:    "12345",
					UserName:     mockLiveVideo1In.Stremer,
					Title:        mockLiveVideo1In.Title,
					ViewerCount:  10,
					StartedAt:    mockLiveVideo1In.StartedDatetime.Format(time.RFC3339),
					ThumbnailUrl: mockLiveVideo1In.Url.String,
				},
			},
		}

		gomock.InOrder(
			mockTwitchApiClient.EXPECT().
				ListBroadcast(gomock.Eq(listBroadcastOptions)).
				Do(func(params []httpClient.RequestParam) {
					a.Equal(listBroadcastOptions, params)
				}).
				Return(listBroadcastResponse, nil).
				Times(1),
		)

		usecase := NewTwitchUpdateLiveVideosUsecase(
			archiveVideoRepository,
			mockTwitchApiClient,
			now.TimeFunc(),
		)

		err = usecase.Handle(ctx)
		a.NoError(err)

		searchBroadcastId := entity.NewVideoBroadcastId(mockLiveVideo2In.BroadcastId)
		actualArchiveVideo, err := archiveVideoRepository.GetByBroadcastId(ctx, searchBroadcastId)
		a.NoError(err)
		a.Equal(entity.VideoStatusEnded.Int(), actualArchiveVideo.GetStatus().Int())
		a.Equal("2022-01-01T15:00:00Z", actualArchiveVideo.GetEndedDatetime().RFC3339())
	})
	t.Run(`正常系/archive_videosに動画情報が1件あり、
		twitchAPIのbroadcastに存在するId = 1は、statusを変更しない`, func(t *testing.T) {
		postgresql.Truncate([]string{"archive_videos"})

		mockLiveVideoIn := mockInput.NewMockArchiveVideoInput(1, url, platformTwitch, videoStatusStreaming, startedDatetime, nil)
		err := archiveVideoRepository.Create(ctx, mockLiveVideoIn)
		a.NoError(err)

		ctrl := gomock.NewController(t)
		mockTwitchApiClient := testTwitch.NewMockTwitchApiClientInterface(ctrl)

		listBroadcastResponse := &twitch.ListBroadcastResponse{
			List: []twitch.ListBroadcast{
				{
					StreamId:     mockLiveVideoIn.BroadcastId,
					UserId:       "12345",
					UserLogin:    "12345",
					UserName:     mockLiveVideoIn.Stremer,
					Title:        mockLiveVideoIn.Title,
					ViewerCount:  10,
					StartedAt:    mockLiveVideoIn.StartedDatetime.Format(time.RFC3339),
					ThumbnailUrl: mockLiveVideoIn.Url.String,
				},
			},
		}

		gomock.InOrder(
			mockTwitchApiClient.EXPECT().
				ListBroadcast(gomock.Eq(listBroadcastOptions)).
				Do(func(params []httpClient.RequestParam) {
					a.Equal(listBroadcastOptions, params)
				}).
				Return(listBroadcastResponse, nil).
				Times(1),
		)

		usecase := NewTwitchUpdateLiveVideosUsecase(
			archiveVideoRepository,
			mockTwitchApiClient,
			now.TimeFunc(),
		)

		err = usecase.Handle(ctx)
		a.NoError(err)

		searchBroadcastId := entity.NewVideoBroadcastId(mockLiveVideoIn.BroadcastId)
		actualArchiveVideo, err := archiveVideoRepository.GetByBroadcastId(ctx, searchBroadcastId)
		a.NoError(err)
		a.Equal(entity.VideoStatusStreaming.Int(), actualArchiveVideo.GetStatus().Int())
		a.Nil(actualArchiveVideo.GetEndedDatetime())
	})
}
