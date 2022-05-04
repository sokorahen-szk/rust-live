package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewArchiveVideo(t *testing.T) {
	a := assert.New(t)

	id := NewVideoId(1)
	broadcastId := NewVideoBroadcastId("xdgdggrh23")
	title := NewVideoTitle("title")
	url := NewVideoUrl("url")
	stremer := NewVideoStremer("stremer")
	thumbnailImage := NewThumbnailImage("src/image.jpg")
	startedDatetime := NewStartedDatetime("2022-04-01T00:00:00Z")
	endedDatetime := NewEndedDatetime("2022-04-01T00:00:00Z")

	archiveVideo := NewArchiveVideo(id, broadcastId, title, url, stremer, thumbnailImage, startedDatetime, endedDatetime)
	a.Equal(id, archiveVideo.GetId())
	a.Equal(broadcastId, archiveVideo.GetBroadcastId())
	a.Equal(title, archiveVideo.GetTitle())
	a.Equal(url, archiveVideo.GetUrl())
	a.Equal(stremer, archiveVideo.GetStremer())
	a.Equal(thumbnailImage, archiveVideo.GetThumbnailImage())
	a.Equal(startedDatetime, archiveVideo.GetStartedDatetime())
	a.Equal(endedDatetime, archiveVideo.GetEndedDatetime())
}
