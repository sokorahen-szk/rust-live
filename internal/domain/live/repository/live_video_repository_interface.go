package repository

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
)

type LiveRepositoryInterface interface {
	List(context.Context) ([]*entity.LiveVideo, error)
}
