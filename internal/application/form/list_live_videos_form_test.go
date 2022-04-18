package form

import (
	"testing"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
)

func Test_NewListLiveVideosForm(t *testing.T) {
	t.Run("検索キーワード", func(t *testing.T) {
		tests := []struct {
			name    string
			arg     string
			isError bool
		}{
			{"検索ワード入力", "あいうえおかきくけこさしすせそabcde", false},
			{"未入力", "", false},
			{"max=20を超過する", "あいうえおかきくけこさしすせそabcdef", true},
		}
		for idx, p := range tests {
			req := &pb.ListLiveVideosRequest{
				SearchKeywords: p.arg,
			}

			t.Run(p.name, func(t *testing.T) {
				form := NewListLiveVideosForm(req)
				err := Validate(form)
				if (err != nil) != p.isError {
					t.Errorf("pattern %d: want %t, name = %s", idx, p.isError, p.name)
					t.Failed()
				}
			})
		}
	})
}
