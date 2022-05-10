package application

import (
	"context"
	"time"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	usecaseBatch "github.com/sokorahen-szk/rust-live/internal/usecase/batch"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
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
	return fetchLiveVideosUsecase{
		liveVideoRepository:    liveVideoRepository,
		archiveVideoRepository: archiveVideoRepository,
		now:                    now,
	}
}

func (usecase updateLiveVideosUsecase) Handle(ctx context.Context) error {
	listArchiveVideos, err := usecase.listArchiveVideos(ctx)
	if err != nil {
		return err
	}

	listLiveVideos, err := usecase.listLiveVideos(ctx)
	if err != nil {
		return err
	}

	videoIds := usecase.filteredEndedVideoIds(listArchiveVideos, listLiveVideos)
	err = usecase.updateVideoStatus(ctx, videoIds)
	if err != nil {
		return err
	}

	return nil
}

func (usecase updateLiveVideosUsecase) listArchiveVideos(ctx context.Context) ([]*entity.ArchiveVideo, error) {
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
			if archiveVideo.GetId() == liveVideo.GetId() {
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

func (usecase updateLiveVideosUsecase) updateVideoStatus(ctx context.Context, videoIds []*entity.VideoId) error {
	updateInput := &input.UpdateArchiveVideoInput{
		Status: entity.NewVideoStatus(entity.VideoStatusEnded),
	}

	for _, videoId := range videoIds {
		err := usecase.archiveVideoRepository.Update(ctx, videoId, updateInput)
		if err != nil {
			return err
		}
	}

	return nil
}
