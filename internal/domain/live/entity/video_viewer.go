package entity

type VideoViewer int

func NewVideoViewer(value int) *VideoViewer {
	if value < 0 {
		panic("入力された値が無効")
	}

	m := VideoViewer(value)
	return &m
}

func (ins VideoViewer) Int() int {
	return int(ins)
}

func (ins VideoViewer) Int32() int32 {
	return int32(ins)
}
