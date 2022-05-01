package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewArchiveVideo(t *testing.T) {
	a := assert.New(t)

	id := NewVideoId(1)
	title := NewVideoTitle("title")
	url := NewVideoUrl("url")
	stremer := NewVideoStremer("stremer")
	thumbnailImage := NewThumbnailImage("src/image.jpg")
	startedDatetime := NewStartedDatetime("2022-04-01T00:00:00Z")
	endedDatetime := NewEndedDatetime("2022-04-01T00:00:00Z")

	archiveVideo := NewArchiveVideo(id, title, url, stremer, thumbnailImage, startedDatetime, endedDatetime)
	a.Equal(id, archiveVideo.Id())
	a.Equal(title, archiveVideo.Title())
	a.Equal(url, archiveVideo.Url())
	a.Equal(stremer, archiveVideo.Stremer())
	a.Equal(thumbnailImage, archiveVideo.ThumbnailImage())
	a.Equal(startedDatetime, archiveVideo.StartedDatetime())
	a.Equal(endedDatetime, archiveVideo.EndedDatetime())
}
