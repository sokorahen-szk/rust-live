package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewThumbnailImage(t *testing.T) {
	a := assert.New(t)

	thumbnailImage := NewThumbnailImage("src/{width}-{height}-image.jpg")
	a.Equal("src/320-320-image.jpg", thumbnailImage.String())
}
