package application

import (
	"testing"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/stretchr/testify/assert"
)

func Test_ToGrpcPagination(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		name string
		arg  *pb.Pagination
	}{
		{
			"リミット=10、レコード=20件、現ページ=1、全ページ=2、前ページ=1、次ページ=2",
			&pb.Pagination{
				Limit:      10,
				Page:       1,
				Prev:       1,
				Next:       2,
				TotalPage:  2,
				TotalCount: 20,
			},
		},
		{
			"リミット=10、レコード=21件、現ページ=2、全ページ=3、前ページ=1、次ページ=3",
			&pb.Pagination{
				Limit:      10,
				Page:       2,
				Prev:       1,
				Next:       3,
				TotalPage:  3,
				TotalCount: 21,
			},
		},
		{
			"リミット=10、レコード=14件、現ページ=14、全ページ=14、前ページ=13、次ページ=14",
			&pb.Pagination{
				Limit:      1,
				Page:       14,
				Prev:       13,
				Next:       14,
				TotalPage:  14,
				TotalCount: 14,
			},
		},
		{
			"リミット=30、レコード=30件、現ページ=1、全ページ=1、前ページ=1、次ページ=1",
			&pb.Pagination{
				Limit:      30,
				Page:       1,
				Prev:       1,
				Next:       1,
				TotalPage:  1,
				TotalCount: 30,
			},
		},
	}
	for _, p := range tests {
		t.Run(p.name, func(t *testing.T) {
			actual := ToGrpcPagination(
				int(p.arg.Page),
				int(p.arg.Limit),
				int(p.arg.TotalCount),
			)
			a.Equal(p.arg, actual)
		})
	}
}
