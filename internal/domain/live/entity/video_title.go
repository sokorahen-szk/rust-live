package entity

type VideoTitle string

func NewVideoTitle(value string) *VideoTitle {
	m := VideoTitle(value)
	return &m
}

func (ins VideoTitle) String() string {
	return string(ins)
}
