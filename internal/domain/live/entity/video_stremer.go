package entity

type VideoStremer string

func NewVideoStremer(value string) *VideoStremer {
	m := VideoStremer(value)
	return &m
}

func (ins VideoStremer) String() string {
	return string(ins)
}
