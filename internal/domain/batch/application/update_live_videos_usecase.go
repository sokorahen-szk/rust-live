package application

import (
	"context"
	"time"

	"github.com/sokorahen-szk/rust-live/internal/domain/common"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	usecaseBatch "github.com/sokorahen-szk/rust-live/internal/usecase/batch"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
	"github.com/sokorahen-szk/rust-live/pkg/logger"
)

type updateLiveVideosUsecase struct {
	liveVideoRepository    repository.LiveVideoRepositoryInterface
	archiveVideoRepository repository.ArchiveVideoRepositoryInterface
	now                    func() time.Time
}

func NewUpdateLiveVideosUsecase(
	liveVideoRepository repository.LiveVideoRepositoryInterface,
	archiveVideoRepository repository.ArchiveVideoRepositoryInterface,
	now func() time.Time,
) usecaseBatch.UpdateLiveVideosUsecaseInterface {
	return updateLiveVideosUsecase{
		liveVideoRepository:    liveVideoRepository,
		archiveVideoRepository: archiveVideoRepository,
		now:                    now,
	}
}

func (usecase updateLiveVideosUsecase) Handle(ctx context.Context) error {
	logger.Info("start update live videos batch")
	defer logger.Info("end update live videos batch")

	now := usecase.now()
	currentDatetime := common.NewDatetimeFromTime(&now)

	listStreamingVideos, err := usecase.listStreamingVideos(ctx)
	if err != nil {
		return err
	}

	listLiveVideos, err := usecase.listLiveVideos(ctx)
	if err != nil {
		return err
	}

	videoIds := usecase.filteredEndedVideoIds(listStreamingVideos, listLiveVideos)
	err = usecase.updateVideoStatus(ctx, videoIds, currentDatetime)
	if err != nil {
		return err
	}

	return nil
}

func (usecase updateLiveVideosUsecase) listStreamingVideos(ctx context.Context) ([]*entity.ArchiveVideo, error) {
	listArchiveVideoInput := &input.ListArchiveVideoInput{
		VideoStatuses: []entity.VideoStatus{entity.VideoStatusStreaming},
	}
	archiveVideos, err := usecase.archiveVideoRepository.List(ctx, listArchiveVideoInput)
	if err != nil {
		return nil, err
	}

	return archiveVideos, nil
}

func (usecase updateLiveVideosUsecase) listLiveVideos(ctx context.Context) ([]*entity.LiveVideo, error) {
	liveVideos, err := usecase.liveVideoRepository.List(ctx, &list.ListLiveVideoInput{})
	if err != nil {
		return nil, err
	}

	return liveVideos, nil
}

func (usecase updateLiveVideosUsecase) filteredEndedVideoIds(archiveVideos []*entity.ArchiveVideo, liveVideos []*entity.LiveVideo) []*entity.VideoId {
	var filteredVideoIds []*entity.VideoId
	for _, archiveVideo := range archiveVideos {
		isLive := false
		for _, liveVideo := range liveVideos {
			if *archiveVideo.GetId() == *liveVideo.GetId() {
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

func (usecase updateLiveVideosUsecase) updateVideoStatus(ctx context.Context, videoIds []*entity.VideoId,
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
