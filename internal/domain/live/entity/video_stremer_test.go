package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewVideoStremer(t *testing.T) {
	a := assert.New(t)

	videoStremer := NewVideoStremer("video_stremer")
	a.Equal("video_stremer", videoStremer.String())
}
