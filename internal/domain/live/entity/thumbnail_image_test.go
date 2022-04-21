package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewThumbnailImage(t *testing.T) {
	a := assert.New(t)

	thumbnailImage := NewThumbnailImage("src/image.jpg")
	a.Equal("src/image.jpg", thumbnailImage.String())
}
