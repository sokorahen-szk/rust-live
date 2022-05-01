package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewLiveVideo(t *testing.T) {
	a := assert.New(t)

	id := NewVideoId(1)
	title := NewVideoTitle("title")
	url := NewVideoUrl("url")
	stremer := NewVideoStremer("stremer")
	viewer := NewVideoViewer(12)
	thumbnailImage := NewThumbnailImage("src/image.jpg")
	startedDatetime := NewStartedDatetime("2022-04-01T00:00:00Z")
	elapsedTimes := NewElapsedTimes(60)

	archiveVideo := NewLiveVideo(id, title, url, stremer, viewer, thumbnailImage, startedDatetime, elapsedTimes)
	a.Equal(id, archiveVideo.GetId())
	a.Equal(title, archiveVideo.GetTitle())
	a.Equal(url, archiveVideo.GetUrl())
	a.Equal(stremer, archiveVideo.GetStremer())
	a.Equal(viewer, archiveVideo.GetViewer())
	a.Equal(thumbnailImage, archiveVideo.GetThumbnailImage())
	a.Equal(startedDatetime, archiveVideo.GetStartedDatetime())
	a.Equal(elapsedTimes, archiveVideo.GetElapsedTimes())
}
