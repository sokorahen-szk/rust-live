package entity

import "fmt"

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

func (ins ElapsedTimes) Hours() int {
	return int(ins.Int() / 3600)
}

func (ins ElapsedTimes) Minutes() int {
	return int((ins.Int() / 60) % 60)
}

func (ins ElapsedTimes) Seconds() int {
	return int(ins.Int() % 60)
}

func (ins ElapsedTimes) Text() string {
	if ins.Int() < 60 {
		return fmt.Sprintf("%02ds", ins.Seconds())
	}
	if ins.Int() < 3600 {
		return fmt.Sprintf("%02dm%02ds", ins.Minutes(), ins.Seconds())
	}

	return fmt.Sprintf(
		"%02dh%02dm%02ds",
		ins.Hours(),
		ins.Minutes(),
		ins.Seconds(),
	)
}
