package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewVideoViewer(t *testing.T) {
	a := assert.New(t)

	videoViewer := NewVideoViewer(1)
	a.Equal(1, videoViewer.Int())
}
