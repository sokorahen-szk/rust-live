package entity

type LiveVideoSortKey int

const (
	LiveVideoViewerAsc           LiveVideoSortKey = 1
	LiveVideoViewerDesc          LiveVideoSortKey = 2
	LiveVideoStartedDatetimeAsc  LiveVideoSortKey = 3
	LiveVideoStartedDatetimeDesc LiveVideoSortKey = 4
)

func NewLiveVideoSortKey(key int) *LiveVideoSortKey {
	m := LiveVideoSortKey(key)
	return &m
}

func (ins LiveVideoSortKey) Int() int {
	return int(ins)
}

func (ins LiveVideoSortKey) IsSortKey() bool {
	return int(ins) > 0 && int(ins) < 5
}
