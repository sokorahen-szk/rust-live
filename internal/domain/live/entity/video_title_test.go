package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewVideoTitle(t *testing.T) {
	a := assert.New(t)

	videoTitle := NewVideoTitle("video_title")
	a.Equal("video_title", videoTitle.String())
}
