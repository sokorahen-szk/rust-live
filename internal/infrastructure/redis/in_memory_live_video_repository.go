package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
)

type inMemoryLiveVideoRepository struct{}

func NewInMemoryLiveVideoRepository() repository.LiveVideoRepositoryInterface {
	return inMemoryLiveVideoRepository{}
}

func (repository inMemoryLiveVideoRepository) List(ctx context.Context, input *list.ListLiveVideosInput) ([]*entity.LiveVideo, error) {
	if ctx.Value("error") == "error" {
		return nil, errors.New("error")
	}

	return repository.liveVideos(), nil
}

func (repository inMemoryLiveVideoRepository) liveVideos() []*entity.LiveVideo {
	liveVideos := make([]*entity.LiveVideo, 0)
	for i := 1; i <= 10; i++ {
		liveVideo := entity.NewLiveVideo(
			entity.NewVideoId(i),
			entity.NewVideoTitle("test"),
			entity.NewVideoStremer(fmt.Sprintf("stremer%d", i)),
			entity.NewThumbnailImage(fmt.Sprintf("src/image%d.jpg", i)),
			entity.NewStartedDatetime("2022-04-01T00:00:00Z"),
			entity.NewElapsedTimes(1),
		)
		liveVideos = append(liveVideos, liveVideo)
	}

	return liveVideos
}
