package entity

type LiveVideoSortKey int

const (
	LiveVideoViewerAsc           LiveVideoSortKey = 1
	LiveVideoViewerDesc          LiveVideoSortKey = 2
	LiveVideoStartedDatetimeAsc  LiveVideoSortKey = 3
	LiveVideoStartedDatetimeDesc LiveVideoSortKey = 4
)

var liveVideoSortKeyValues = []LiveVideoSortKey{
	LiveVideoViewerAsc,
	LiveVideoViewerDesc,
	LiveVideoStartedDatetimeAsc,
	LiveVideoStartedDatetimeDesc,
}

func NewLiveVideoSortKey(value LiveVideoSortKey) *LiveVideoSortKey {
	m := LiveVideoSortKey(value)
	return &m
}

func NewLiveVideoSortKeyFromInt(n int) *LiveVideoSortKey {
	for _, p := range liveVideoSortKeyValues {
		if int(p) == n {
			return &p
		}
	}

	return nil
}

func (ins LiveVideoSortKey) Int() int {
	return int(ins)
}

func (ins LiveVideoSortKey) IsSortKey() bool {
	return int(ins) > 0 && int(ins) < 5
}
