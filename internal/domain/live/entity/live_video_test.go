package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewLiveVideo(t *testing.T) {
	a := assert.New(t)

	id := NewVideoId(1)
	broadcastId := NewVideoBroadcastId("xdgdggrh23")
	title := NewVideoTitle("title")
	url := NewVideoUrl("url")
	stremer := NewVideoStremer("stremer")
	viewer := NewVideoViewer(12)
	platform := NewPlatform(PlatformTwitch)
	thumbnailImage := NewThumbnailImage("src/image.jpg")
	startedDatetime := NewStartedDatetime("2022-04-01T00:00:00Z")
	elapsedTimes := NewElapsedTimes(60)

	archiveVideo := NewLiveVideo(id, broadcastId, title, url, stremer, viewer, platform, thumbnailImage, startedDatetime, elapsedTimes)
	a.Equal(id, archiveVideo.GetId())
	a.Equal(broadcastId, archiveVideo.GetBroadcastId())
	a.Equal(title, archiveVideo.GetTitle())
	a.Equal(url, archiveVideo.GetUrl())
	a.Equal(stremer, archiveVideo.GetStremer())
	a.Equal(viewer, archiveVideo.GetViewer())
	a.Equal(platform, archiveVideo.GetPlatform())
	a.Equal(thumbnailImage, archiveVideo.GetThumbnailImage())
	a.Equal(startedDatetime, archiveVideo.GetStartedDatetime())
	a.Equal(elapsedTimes, archiveVideo.GetElapsedTimes())
}
