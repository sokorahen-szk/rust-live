package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLiveVideoSortKey(t *testing.T) {
	tests := []struct {
		name   string
		arg    int
		result bool
	}{
		{"ソートキー未設定", 0, false},
		{"LiveVideoViewerAsc", 1, true},
		{"LiveVideoViewerDesc", 2, true},
		{"LiveVideoStartedDatetimeAsc", 3, true},
		{"LiveVideoStartedDatetimeDesc", 4, true},
		{"不明なソートキー設定", 5, false},
	}
	for idx, p := range tests {
		t.Run(p.name, func(t *testing.T) {
			videoSortKey := NewLiveVideoSortKey(p.arg)
			if videoSortKey.IsSortKey() != p.result {
				t.Errorf("pattern %d: want %t, name = %s", idx, p.result, p.name)
				t.Failed()
			}

			assert.Equal(t, p.arg, videoSortKey.Int())
		})
	}
}
