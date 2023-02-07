package application_batch

import (
	"context"
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

type twitchUpdateLiveVideosUsecase struct {
	archiveVideoRepository repository.ArchiveVideoRepositoryInterface
	twitchApiClient        twitch.TwitchApiClientInterface
	now                    func() time.Time
}

func NewTwitchUpdateLiveVideosUsecase(
	archiveVideoRepository repository.ArchiveVideoRepositoryInterface,
	twitchApiClient twitch.TwitchApiClientInterface,
	now func() time.Time,
) usecaseBatch.TwitchUpdateLiveVideosUsecaseInterface {
	return twitchUpdateLiveVideosUsecase{
		archiveVideoRepository: archiveVideoRepository,
		twitchApiClient:        twitchApiClient,
		now:                    now,
	}
}

func (usecase twitchUpdateLiveVideosUsecase) Handle(ctx context.Context) error {
	now := usecase.now()
	currentDatetime := common.NewDatetimeFromTime(&now)

	logger.Info(fmt.Sprintf("start update live videos batch %s", now.Format(time.RFC3339)))

	listStreamingVideos, err := usecase.listStreamingVideos(ctx)
	if err != nil {
		return err
	}

	listBroadcastResponse, err := usecase.listTwitchBroadcast()
	if err != nil {
		return err
	}

	videoIds := usecase.filteredEndedVideoIds(listStreamingVideos, listBroadcastResponse)
	err = usecase.updateVideoStatus(ctx, videoIds, currentDatetime)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("end update live videos batch %s", now.Format(time.RFC3339)))
	return nil
}

func (usecase twitchUpdateLiveVideosUsecase) listStreamingVideos(ctx context.Context) ([]*entity.ArchiveVideo, error) {
	listArchiveVideoInput := &input.ListArchiveVideoInput{
		VideoStatuses: []entity.VideoStatus{entity.VideoStatusStreaming},
	}
	archiveVideos, err := usecase.archiveVideoRepository.List(ctx, listArchiveVideoInput)
	if err != nil {
		return nil, err
	}

	return archiveVideos, nil
}

func (usecase twitchUpdateLiveVideosUsecase) listTwitchBroadcast() (*twitch.ListBroadcastResponse, error) {
	options := []httpClient.RequestParam{
		{Key: "language", Value: "ja"},
		{Key: "game_id", Value: twitch.RustGameId},
		{Key: "type", Value: "live"},
		{Key: "first", Value: "100"},
	}

	return usecase.twitchApiClient.ListBroadcast(options)
}

func (usecase twitchUpdateLiveVideosUsecase) filteredEndedVideoIds(archiveVideos []*entity.ArchiveVideo,
	listBroadcastResponse *twitch.ListBroadcastResponse) []*entity.VideoId {

	var filteredVideoIds []*entity.VideoId
	for _, archiveVideo := range archiveVideos {
		isLive := false
		for _, broadcast := range listBroadcastResponse.List {
			startedDatetime := entity.NewStartedDatetime(broadcast.StartedAt)
			if archiveVideo.GetStartedDatetime().Equal(startedDatetime.Time()) &&
				archiveVideo.GetStremer().String() == broadcast.UserName {
				isLive = true
				break
			}
		}

		if isLive {
			continue
		}

		filteredVideoIds = append(filteredVideoIds, archiveVideo.GetId())
	}

	return filteredVideoIds
}

func (usecase twitchUpdateLiveVideosUsecase) updateVideoStatus(ctx context.Context, videoIds []*entity.VideoId,
	currentDatetime *common.Datetime) error {
	updateInput := &input.UpdateArchiveVideoInput{
		Status:        entity.NewVideoStatus(entity.VideoStatusEnded),
		EndedDatetime: entity.NewEndedDatetimeFromTime(currentDatetime.Time()),
	}

	for _, videoId := range videoIds {
		_, err := usecase.archiveVideoRepository.Update(ctx, videoId, updateInput)
		if err != nil {
			return err
		}
	}

	return nil
}
