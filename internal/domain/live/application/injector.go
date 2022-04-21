package application

import (
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
)

func NewInjectListLiveVideosUsecase() list.ListLiveVideosUsecaseInterface {
	return NewListLiveVideosUsecase()
}
