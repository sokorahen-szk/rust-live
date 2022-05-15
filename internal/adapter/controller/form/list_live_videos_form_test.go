package form

import (
	"testing"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/stretchr/testify/assert"
)

func Test_NewListLiveVideosForm(t *testing.T) {

	const (
		correctSearchKeywords string                        = "example"
		correctSort           pb.ListLiveVideosRequest_Sort = pb.ListLiveVideosRequest_SORT_UNKNOWN
		correctPage           int32                         = 1
		correctLimit          int32                         = 10
	)

	t.Run("検索キーワード", func(t *testing.T) {
		tests := []struct {
			name    string
			arg     string
			isError bool
		}{
			{"未設定", "", false},
			{"検索ワード入力", "あいうえおかきくけこさしすせそabcde", false},
			{"max=20を超過する", "あいうえおかきくけこさしすせそabcdef", true},
		}
		for idx, p := range tests {
			req := &pb.ListLiveVideosRequest{
				SearchKeywords: p.arg,
				Sort:           correctSort,
				Page:           correctPage,
				Limit:          correctLimit,
			}

			t.Run(p.name, func(t *testing.T) {
				form := NewListLiveVideosForm(req)
				err := Validate(form)
				if (err != nil) != p.isError {
					t.Errorf("pattern %d: want %t, name = %s, err: %s", idx, p.isError, p.name, err)
					t.Failed()
				}
			})
		}
	})

	t.Run("ソートキー", func(t *testing.T) {
		tests := []struct {
			name    string
			arg     pb.ListLiveVideosRequest_Sort
			isError bool
		}{
			{"未設定", pb.ListLiveVideosRequest_SORT_UNKNOWN, false},
			{"SORT_PLATFORM", pb.ListLiveVideosRequest_SORT_PLATFORM, false},
			{"SORT_VIEWER_ASC", pb.ListLiveVideosRequest_SORT_VIEWER_ASC, false},
			{"SORT_VIEWER_DESC", pb.ListLiveVideosRequest_SORT_VIEWER_DESC, false},
			{"SORT_STARTED_DATETIME_ASC", pb.ListLiveVideosRequest_SORT_STARTED_DATETIME_ASC, false},
			{"SORT_STARTED_DATETIME_DESC", pb.ListLiveVideosRequest_SORT_STARTED_DATETIME_DESC, false},
		}
		for idx, p := range tests {
			req := &pb.ListLiveVideosRequest{
				SearchKeywords: correctSearchKeywords,
				Sort:           p.arg,
				Page:           correctPage,
				Limit:          correctLimit,
			}

			t.Run(p.name, func(t *testing.T) {
				form := NewListLiveVideosForm(req)
				err := Validate(form)
				if (err != nil) != p.isError {
					t.Errorf("pattern %d: want %t, name = %s, err: %s", idx, p.isError, p.name, err)
					t.Failed()
				}
			})
		}
	})
	t.Run("現在ページ", func(t *testing.T) {
		tests := []struct {
			name    string
			arg     int32
			isError bool
		}{
			{"未設定", 0, true},
			{"規定値", 1, false},
		}
		for idx, p := range tests {
			req := &pb.ListLiveVideosRequest{
				SearchKeywords: correctSearchKeywords,
				Sort:           correctSort,
				Page:           p.arg,
				Limit:          correctLimit,
			}

			t.Run(p.name, func(t *testing.T) {
				form := NewListLiveVideosForm(req)
				err := Validate(form)
				if (err != nil) != p.isError {
					t.Errorf("pattern %d: want %t, name = %s, err: %s", idx, p.isError, p.name, err)
					t.Failed()
				}
			})
		}
	})
	t.Run("取得する件数", func(t *testing.T) {
		tests := []struct {
			name    string
			arg     int32
			isError bool
		}{
			{"未設定", 0, false},
			{"規定値", 1, false},
		}
		for idx, p := range tests {
			req := &pb.ListLiveVideosRequest{
				SearchKeywords: correctSearchKeywords,
				Sort:           correctSort,
				Page:           correctPage,
				Limit:          p.arg,
			}

			t.Run(p.name, func(t *testing.T) {
				form := NewListLiveVideosForm(req)
				err := Validate(form)
				if (err != nil) != p.isError {
					t.Errorf("pattern %d: want %t, name = %s, err: %s", idx, p.isError, p.name, err)
					t.Failed()
				}
			})
		}
	})

	t.Run("Getter値チェック", func(t *testing.T) {
		req := &pb.ListLiveVideosRequest{
			SearchKeywords: correctSearchKeywords,
			Sort:           correctSort,
			Page:           correctPage,
			Limit:          correctLimit,
		}

		form := NewListLiveVideosForm(req)
		err := Validate(form)
		assert.NoError(t, err)
		assert.Equal(t, correctSearchKeywords, form.GetSearchKeywords())
		assert.Equal(t, int(correctSort.Number()), form.GetSort())
		assert.Equal(t, int(correctPage), form.GetPage())
		assert.Equal(t, int(correctLimit), form.GetLimit())
	})
}
