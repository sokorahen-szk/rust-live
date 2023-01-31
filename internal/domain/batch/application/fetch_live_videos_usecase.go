package application

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sokorahen-szk/rust-live/internal/domain/common"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch/twitch"
	usecaseBatch "github.com/sokorahen-szk/rust-live/internal/usecase/batch"
	httpClient "github.com/sokorahen-szk/rust-live/pkg/http"
	"github.com/sokorahen-szk/rust-live/pkg/logger"
)

type fetchLiveVideosUsecase struct {
	liveVideoRepository    repository.LiveVideoRepositoryInterface
	archiveVideoRepository repository.ArchiveVideoRepositoryInterface
	twitchApiClient        twitch.TwitchApiClientInterface
	now                    func() time.Time
}

func NewFetchLiveVideosUsecase(
	liveVideoRepository repository.LiveVideoRepositoryInterface,
	archiveVideoRepository repository.ArchiveVideoRepositoryInterface,
	twitchApiClient twitch.TwitchApiClientInterface,
	now func() time.Time,
) usecaseBatch.FetchLiveVideosUsecaseInterface {
	return fetchLiveVideosUsecase{
		liveVideoRepository:    liveVideoRepository,
		archiveVideoRepository: archiveVideoRepository,
		twitchApiClient:        twitchApiClient,
		now:                    now,
	}
}

func (usecase fetchLiveVideosUsecase) Handle(ctx context.Context) error {
	now := usecase.now()
	currentDatetime := common.NewDatetimeFromTime(&now)

	logger.Info(fmt.Sprintf("start fetch live videos batch %s", now.Format(time.RFC3339)))

	liveVideos, err := usecase.fetchTwitchApiDataToLocalStorage(ctx, currentDatetime)
	if err != nil {
		return err
	}
	err = usecase.createLiveVideos(ctx, liveVideos)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("end fetch live videos batch %s", now.Format(time.RFC3339)))
	return nil
}

func (usecase fetchLiveVideosUsecase) listTwitchBroadcast() (*twitch.ListBroadcastResponse, error) {
	options := []httpClient.RequestParam{
		{Key: "language", Value: "ja"},
		{Key: "game_id", Value: twitch.RustGameId},
		{Key: "type", Value: "live"},
		{Key: "first", Value: "100"},
	}

	return usecase.twitchApiClient.ListBroadcast(options)
}

func (usecase fetchLiveVideosUsecase) listTwitchVideoByUserId(userId string) (*twitch.ListVideoByUserIdResponse, error) {
	options := []httpClient.RequestParam{
		{Key: "first", Value: "1"},
	}

	return usecase.twitchApiClient.ListVideoByUserId(userId, options)
}

func (usecase fetchLiveVideosUsecase) createArchiveVideos(ctx context.Context, in *input.ArchiveVideoInput) *entity.VideoId {
	searchBroadcastId := entity.NewVideoBroadcastId(in.BroadcastId)
	archiveVideo, _ := usecase.archiveVideoRepository.GetByBroadcastId(ctx, searchBroadcastId)

	if archiveVideo != nil {
		return archiveVideo.GetId()
	}

	err := usecase.archiveVideoRepository.Create(ctx, in)
	if err != nil {
		logger.Errorf("failed archiveVideoRepository.Create() err: %+v", err)
		return nil
	}

	videoId := entity.NewVideoId(in.Id)
	return videoId
}

func (usecase fetchLiveVideosUsecase) createLiveVideos(ctx context.Context, liveVideos []*entity.LiveVideo) error {
	return usecase.liveVideoRepository.Create(ctx, liveVideos)
}

func (usecase fetchLiveVideosUsecase) fetchTwitchApiDataToLocalStorage(ctx context.Context,
	currentDatetime *common.Datetime) ([]*entity.LiveVideo, error) {
	ListBroadcastResponse, err := usecase.listTwitchBroadcast()
	if err != nil {
		return nil, err
	}

	liveVideos := make([]*entity.LiveVideo, 0)
	for _, broadcastData := range ListBroadcastResponse.List {
		listVideoByUserIdRes, err := usecase.listTwitchVideoByUserId(broadcastData.UserId)
		if err != nil {
			logger.Infof("failed twitchApiClient.ListBroadcast() err: %+v", err)
			continue
		}

		var broadcastId *entity.VideoBroadcastId
		var archiveVideoUrl *sql.NullString
		if len(listVideoByUserIdRes.List) > 0 {
			broadcastId = entity.NewVideoBroadcastId(listVideoByUserIdRes.List[0].Id)
			videoUrl := usecase.convertTwtichArchiveVideoUrl(broadcastId)
			archiveVideoUrl = &sql.NullString{String: videoUrl.String(), Valid: true}
		} else {
			broadcastId = entity.NewVideoBroadcastId(broadcastData.StreamId)
			archiveVideoUrl = &sql.NullString{String: "", Valid: false}
		}

		liveVideoUrl := usecase.convertTwitchLiveVideoUrl(broadcastData.UserLogin)
		stremer := entity.NewVideoStremer(broadcastData.UserName)
		title := entity.NewVideoTitle(broadcastData.Title)
		viewer := entity.NewVideoViewer(broadcastData.ViewerCount)
		platform := entity.NewPlatform(entity.PlatformTwitch)
		thumbnailImage := entity.NewThumbnailImage(broadcastData.ThumbnailUrl)
		startedDatetime := entity.NewStartedDatetime(broadcastData.StartedAt)
		elapsedTimes := entity.NewElapsedTimes(currentDatetime.DiffSeconds(startedDatetime.Time()))

		in := &input.ArchiveVideoInput{
			BroadcastId:     broadcastId.String(),
			Title:           title.String(),
			Url:             archiveVideoUrl,
			Stremer:         stremer.String(),
			Platform:        platform.Int(),
			Status:          entity.VideoStatusStreaming.Int(),
			ThumbnailImage:  thumbnailImage.String(),
			StartedDatetime: startedDatetime.Time(),
		}
		videoId := usecase.createArchiveVideos(ctx, in)
		if videoId == nil {
			continue
		}

		liveVideos = append(liveVideos, entity.NewLiveVideo(
			videoId,
			broadcastId,
			title,
			liveVideoUrl,
			stremer,
			viewer,
			platform,
			thumbnailImage,
			startedDatetime,
			elapsedTimes,
		))
	}

	return liveVideos, nil
}

func (usecase fetchLiveVideosUsecase) convertTwtichArchiveVideoUrl(broadcastId *entity.VideoBroadcastId) *entity.VideoUrl {
	return entity.NewVideoUrl(fmt.Sprintf("https://www.twitch.tv/videos/%s", broadcastId.String()))
}

func (usecase fetchLiveVideosUsecase) convertTwitchLiveVideoUrl(userLoginName string) *entity.VideoUrl {
	return entity.NewVideoUrl(fmt.Sprintf("https://www.twitch.tv/%s", userLoginName))
}
