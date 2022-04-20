package live

type ArchiveVideoId int

func NewArchiveVideoId(value int) *ArchiveVideoId {

	if value == 0 {
		panic("入力された値が無効")
	}

	m := ArchiveVideoId(value)
	return &m
}

func (ins ArchiveVideoId) Int() int {
	return int(ins)
}
