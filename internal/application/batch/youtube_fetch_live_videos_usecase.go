package application_batch

import (
	"context"
	"time"

	"github.com/sokorahen-szk/rust-live/internal/domain/common"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch/youtube"
	usecaseBatch "github.com/sokorahen-szk/rust-live/internal/usecase/batch"
	"github.com/sokorahen-szk/rust-live/pkg/logger"
)

type youtubeFetchLiveVideosUsecase struct {
	liveVideoRepository    repository.LiveVideoRepositoryInterface
	archiveVideoRepository repository.ArchiveVideoRepositoryInterface
	youtubeApiClient       youtube.YouTubeApiClientInterface
	now                    func() time.Time
}

func NewYoutubeFetchLiveVideosUsecase(
	liveVideoRepository repository.LiveVideoRepositoryInterface,
	archiveVideoRepository repository.ArchiveVideoRepositoryInterface,
	youtubeApiClient youtube.YouTubeApiClientInterface,
	now func() time.Time,
) usecaseBatch.YoutubeFetchLiveVideosUsecaseInterface {
	return youtubeFetchLiveVideosUsecase{
		liveVideoRepository:    liveVideoRepository,
		archiveVideoRepository: archiveVideoRepository,
		youtubeApiClient:       youtubeApiClient,
		now:                    now,
	}
}

func (usecase youtubeFetchLiveVideosUsecase) Handle(ctx context.Context) error {
	now := usecase.now()
	currentDatetime := common.NewDatetimeFromTime(&now)

	usecase.fetchTwitchApiDataToLocalStorage(ctx, currentDatetime)

	return nil
}

func (usecase youtubeFetchLiveVideosUsecase) fetchTwitchApiDataToLocalStorage(ctx context.Context,
	currentDatetime *common.Datetime) (interface{}, error) {

	listBroadcastResponse, err := usecase.youtubeApiClient.ListBroadcast()
	if err != nil {
		return nil, err
	}

	logger.Infof("%+v", listBroadcastResponse)

	return nil, nil
}
