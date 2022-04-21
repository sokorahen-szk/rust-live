package entity

type ThumbnailImage string

func NewThumbnailImage(value string) *ThumbnailImage {
	m := ThumbnailImage(value)
	return &m
}

func (ins ThumbnailImage) String() string {
	return string(ins)
}
