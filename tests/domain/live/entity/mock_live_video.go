package mockEntity

import (
	"fmt"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
)

func NewMockLiveVideo(i int) *entity.LiveVideo {
	videoId := entity.NewVideoId(i)
	videoBroadcastId := entity.NewVideoBroadcastId(fmt.Sprintf("gefefgh%d", i))
	videoTitle := entity.NewVideoTitle(fmt.Sprintf("title_%d", i))
	videoUrl := entity.NewVideoUrl(fmt.Sprintf("https://example.com/%d", i))
	videoStremer := entity.NewVideoStremer(fmt.Sprintf("太郎%d", i))
	videoViewer := entity.NewVideoViewer(10)
	platform := entity.NewPlatform(entity.PlatformTwitch)
	thumbnailImage := entity.NewThumbnailImage(fmt.Sprintf("https://example.com/test%d.jpg", i))
	startedDatetime := entity.NewStartedDatetime("2022-12-31T15:00:00Z")
	elapsedTimes := entity.NewElapsedTimes(10000)

	return entity.NewLiveVideo(
		videoId,
		videoBroadcastId,
		videoTitle,
		videoUrl,
		videoStremer,
		videoViewer,
		platform,
		thumbnailImage,
		startedDatetime,
		elapsedTimes,
	)
}
