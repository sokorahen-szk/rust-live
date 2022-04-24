package entity

type VideoId int

func NewVideoId(value int) *VideoId {
	if value == 0 {
		panic("入力された値が無効")
	}

	m := VideoId(value)
	return &m
}

func (ins VideoId) Int() int {
	return int(ins)
}

func (ins VideoId) Int32() int32 {
	return int32(ins)
}
