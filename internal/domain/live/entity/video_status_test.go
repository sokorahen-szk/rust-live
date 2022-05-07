package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewVideoStatus(t *testing.T) {
	a := assert.New(t)

	videoStatus := NewVideoStatus(VideoStatusStreaming)
	a.Equal(1, videoStatus.Int())
	a.Equal("Streaming", videoStatus.String())

	videoStatus = NewVideoStatus(VideoStatusEnded)
	a.Equal(2, videoStatus.Int())
	a.Equal("Ended", videoStatus.String())
}

func Test_NewVideoStatusFromInt(t *testing.T) {
	a := assert.New(t)

	videoStatus := NewVideoStatusFromInt(1)
	a.Equal(1, videoStatus.Int())
	a.Equal("Streaming", videoStatus.String())

	videoStatus = NewVideoStatusFromInt(2)
	a.Equal(2, videoStatus.Int())
	a.Equal("Ended", videoStatus.String())
}
