package application

import (
	"context"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
)

type ListLiveVideosUsecase struct{}

func NewListLiveVideosUsecase() list.ListLiveVideosUsecaseInterface {
	return ListLiveVideosUsecase{}
}

func (ins ListLiveVideosUsecase) Handle(context.Context, *list.ListLiveVideosInput) (*pb.ListLiveVideosResponse, error) {
	return &pb.ListLiveVideosResponse{}, nil
}
