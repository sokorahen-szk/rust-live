package entity

import (
	"testing"
)

func TestNewLiveVideoSortKey(t *testing.T) {
	tests := []struct {
		name   string
		arg    LiveVideoSortKey
		result bool
	}{
		{"ソートキー未設定", 0, false},
		{"LiveVideoViewerAsc", LiveVideoViewerAsc, true},
		{"LiveVideoViewerDesc", LiveVideoViewerDesc, true},
		{"LiveVideoStartedDatetimeAsc", LiveVideoStartedDatetimeAsc, true},
		{"LiveVideoStartedDatetimeDesc", LiveVideoStartedDatetimeDesc, true},
		{"不明なソートキー設定", 5, false},
	}
	for idx, p := range tests {
		t.Run(p.name, func(t *testing.T) {
			videoSortKey := NewLiveVideoSortKey(p.arg)
			if videoSortKey.IsSortKey() != p.result {
				t.Errorf("pattern %d: want %t, name = %s", idx, p.result, p.name)
				t.Failed()
			}
		})
	}
}
