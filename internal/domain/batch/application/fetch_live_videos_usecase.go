package application

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch/twitch"
	usecaseBatch "github.com/sokorahen-szk/rust-live/internal/usecase/batch"
	"github.com/sokorahen-szk/rust-live/pkg/logger"
)

type fetchLiveVideosUsecase struct {
	liveVideoRepository    repository.LiveVideoRepositoryInterface
	archiveVideoRepository repository.ArchiveVideoRepositoryInterface
	twitchApiClient        twitch.TwitchApiClientInterface
}

func NewFetchLiveVideosUsecase(
	liveVideoRepository repository.LiveVideoRepositoryInterface,
	archiveVideoRepository repository.ArchiveVideoRepositoryInterface,
	twitchApiClient twitch.TwitchApiClientInterface,
) usecaseBatch.FetchLiveVideosUsecaseInterface {
	return fetchLiveVideosUsecase{
		liveVideoRepository:    liveVideoRepository,
		archiveVideoRepository: archiveVideoRepository,
		twitchApiClient:        twitchApiClient,
	}
}

func (usecase fetchLiveVideosUsecase) Handle(ctx context.Context) error {
	ListBroadcastResponse, err := usecase.listTwitchBroadcast()
	if err != nil {
		return err
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
		thumbnailImage := entity.NewThumbnailImage(broadcastData.ThumbnailUrl)
		startedDatetime := entity.NewStartedDatetime(broadcastData.StartedAt)
		// とりあえず１を入れてる. 後ほど現在時刻とstartedDatetimeの差分を出す
		elapsedTimes := entity.NewElapsedTimes(1)

		in := &input.ArchiveVideoInput{
			BroadcastId:     broadcastId.String(),
			Title:           title.String(),
			Url:             archiveVideoUrl,
			Stremer:         stremer.String(),
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
			thumbnailImage,
			startedDatetime,
			elapsedTimes,
		))
	}

	err = usecase.createLiveVideos(ctx, liveVideos)
	if err != nil {
		return err
	}

	return nil
}

func (usecase fetchLiveVideosUsecase) listTwitchBroadcast() (*twitch.ListBroadcastResponse, error) {
	options := []batch.RequestParam{
		{Key: "language", Value: "ja"},
		{Key: "game_id", Value: twitch.RustGameId},
		{Key: "type", Value: "live"},
		{Key: "first", Value: "100"},
	}

	return usecase.twitchApiClient.ListBroadcast(options)
}

func (usecase fetchLiveVideosUsecase) listTwitchVideoByUserId(userId string) (*twitch.ListVideoByUserIdResponse, error) {
	options := []batch.RequestParam{
		{Key: "first", Value: "1"},
	}

	return usecase.twitchApiClient.ListVideoByUserId(userId, options)
}

func (usecase fetchLiveVideosUsecase) createArchiveVideos(ctx context.Context, in *input.ArchiveVideoInput) *entity.VideoId {
	err := usecase.archiveVideoRepository.Create(ctx, in)
	if err != nil {
		logger.Infof("failed archiveVideoRepository.Create() err: %+v", err)
		return nil
	}

	videoId := entity.NewVideoId(in.Id)
	return videoId
}

func (usecase fetchLiveVideosUsecase) createLiveVideos(ctx context.Context, liveVideos []*entity.LiveVideo) error {
	return usecase.liveVideoRepository.Create(ctx, liveVideos)
}

func (usecase fetchLiveVideosUsecase) convertTwtichArchiveVideoUrl(broadcastId *entity.VideoBroadcastId) *entity.VideoUrl {
	return entity.NewVideoUrl(fmt.Sprintf("https://www.twitch.tv/videos/%s", broadcastId.String()))
}

func (usecase fetchLiveVideosUsecase) convertTwitchLiveVideoUrl(userLoginName string) *entity.VideoUrl {
	return entity.NewVideoUrl(fmt.Sprintf("https://www.twitch.tv/%s", userLoginName))
}
