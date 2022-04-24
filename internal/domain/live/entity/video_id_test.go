package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewVideoId(t *testing.T) {
	a := assert.New(t)

	videoId := NewVideoId(1)
	a.Equal(1, videoId.Int())
}
