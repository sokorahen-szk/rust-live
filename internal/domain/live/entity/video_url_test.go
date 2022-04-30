package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Url(t *testing.T) {
	a := assert.New(t)

	videoUrl := NewVideoUrl("video_url")
	a.Equal("video_url", videoUrl.String())
}
