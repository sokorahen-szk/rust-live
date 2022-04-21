package list

import (
	"context"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
)

type ListLiveVideosUsecaseInterface interface {
	Handle(context.Context, *ListLiveVideosInput) (*pb.ListLiveVideosResponse, error)
}
