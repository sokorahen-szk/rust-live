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
	thumbnailImage := NewThumbnailImage("src/image.jpg")
	startedDatetime := NewStartedDatetime("2022-04-01T00:00:00Z")
	elapsedTimes := NewElapsedTimes(60)

	archiveVideo := NewLiveVideo(id, title, url, stremer, thumbnailImage, startedDatetime, elapsedTimes)
	a.Equal(id, archiveVideo.GetId())
	a.Equal(title, archiveVideo.GetTitle())
	a.Equal(url, archiveVideo.GetUrl())
	a.Equal(stremer, archiveVideo.GetStremer())
	a.Equal(thumbnailImage, archiveVideo.GetThumbnailImage())
	a.Equal(startedDatetime, archiveVideo.GetStartedDatetime())
	a.Equal(elapsedTimes, archiveVideo.GetElapsedTimes())
}
