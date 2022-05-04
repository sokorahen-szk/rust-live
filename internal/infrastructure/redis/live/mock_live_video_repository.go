package live

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
)

type mockLiveVideoRepository struct{}

func NewMockLiveVideoRepository() repository.LiveVideoRepositoryInterface {
	return mockLiveVideoRepository{}
}

func (repository mockLiveVideoRepository) List(ctx context.Context, input *list.ListLiveVideosInput) ([]*entity.LiveVideo, error) {
	if ctx.Value("error") == "error" {
		return nil, errors.New("error")
	}

	return repository.liveVideos(input.SearchKeywords()), nil
}

func (repository mockLiveVideoRepository) liveVideos(searchKeywords string) []*entity.LiveVideo {
	liveVideos := make([]*entity.LiveVideo, 0)
	for i := 1; i <= 10; i++ {

		videoTitle := entity.NewVideoTitle(fmt.Sprintf("test%d", i))
		videoUrl := entity.NewVideoUrl("https://example.com")
		videoStremer := entity.NewVideoStremer(fmt.Sprintf("stremer%d", i))
		videoBroadcastId := entity.NewVideoBroadcastId(fmt.Sprintf("gefefgh%d", i))

		if !strings.Contains(videoTitle.String(), searchKeywords) &&
			!strings.Contains(videoStremer.String(), searchKeywords) {
			continue
		}

		liveVideo := entity.NewLiveVideo(
			entity.NewVideoId(i),
			videoBroadcastId,
			videoTitle,
			videoUrl,
			videoStremer,
			entity.NewVideoViewer(12),
			entity.NewThumbnailImage(fmt.Sprintf("src/image%d.jpg", i)),
			entity.NewStartedDatetime("2022-04-01T00:00:00Z"),
			entity.NewElapsedTimes(1),
		)
		liveVideos = append(liveVideos, liveVideo)
	}

	return liveVideos
}

func (repository mockLiveVideoRepository) Create(_ context.Context, _ []*entity.LiveVideo) error {
	return nil
}
