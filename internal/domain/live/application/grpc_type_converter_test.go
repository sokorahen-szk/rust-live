package application

import (
	"testing"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/stretchr/testify/assert"
)

func Test_ToGrpcPagination(t *testing.T) {
	a := assert.New(t)

	expected := &pb.Pagination{
		Limit:      10,
		Page:       1,
		Prev:       1,
		Next:       2,
		TotalPage:  2,
		TotalCount: 20,
	}
	actual := ToGrpcPagination(1, 10, 20)
	a.Equal(actual, expected)

	expected = &pb.Pagination{
		Limit:      10,
		Page:       2,
		Prev:       1,
		Next:       3,
		TotalPage:  3,
		TotalCount: 21,
	}
	actual = ToGrpcPagination(2, 10, 21)
	a.Equal(actual, expected)
}
