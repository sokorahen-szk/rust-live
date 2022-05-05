package entity

import (
	"fmt"
	"strings"
)

type ThumbnailImage string

const (
	ThumbnailImageWidth  int = 320
	ThumbnailImageHeight int = 320
)

func NewThumbnailImage(value string) *ThumbnailImage {
	value = strings.Replace(value, "{width}", fmt.Sprintf("%d", ThumbnailImageWidth), -1)
	value = strings.Replace(value, "{height}", fmt.Sprintf("%d", ThumbnailImageHeight), -1)

	m := ThumbnailImage(value)
	return &m
}

func (ins ThumbnailImage) String() string {
	return string(ins)
}
