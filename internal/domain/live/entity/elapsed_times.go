package entity

type ElapsedTimes int

func NewElapsedTimes(value int) *ElapsedTimes {
	if value == 0 {
		panic("入力された値が無効")
	}

	m := ElapsedTimes(value)
	return &m
}

func (ins ElapsedTimes) Int() int {
	return int(ins)
}

func (ins ElapsedTimes) Int32() int32 {
	return int32(ins)
}

func (ins ElapsedTimes) Hour() int {
	return int(ins.Int() / 3600)
}
