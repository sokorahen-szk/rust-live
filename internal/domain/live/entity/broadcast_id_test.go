package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewVideoBroadcastId(t *testing.T) {
	a := assert.New(t)

	videoBroadcastId := NewVideoBroadcastId("video_broadcast_id")
	a.Equal("video_broadcast_id", videoBroadcastId.String())
}
