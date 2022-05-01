package entity

type VideoUrl string

func NewVideoUrl(value string) *VideoUrl {
	m := VideoUrl(value)
	return &m
}

func (ins VideoUrl) String() string {
	return string(ins)
}
