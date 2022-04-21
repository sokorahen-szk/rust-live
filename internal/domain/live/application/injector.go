package application

import (
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"

	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
)

func NewInjectListLiveVideosUsecase() list.ListLiveVideosUsecaseInterface {

	liveVideoRepository := redis.NewInMemoryLiveVideoRepository()
	return NewListLiveVideosUsecase(liveVideoRepository)
}
