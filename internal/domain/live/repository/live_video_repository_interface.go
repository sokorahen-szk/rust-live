package repository

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
)

type LiveVideoRepositoryInterface interface {
	Create(context.Context, []*entity.LiveVideo) error
	List(context.Context, *list.ListLiveVideoInput) ([]*entity.LiveVideo, error)
	Count(context.Context) (int, error)
}
