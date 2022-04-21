package repository

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
)

type LiveVideoRepositoryInterface interface {
	List(context.Context, *list.ListLiveVideosInput) ([]*entity.LiveVideo, error)
}
