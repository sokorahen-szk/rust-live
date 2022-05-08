package repository

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
)

type ArchiveVideoRepositoryInterface interface {
	GetByBroadcastId(context.Context, *entity.VideoBroadcastId) (*entity.ArchiveVideo, error)
	List(context.Context, *input.ListArchiveVideoInput) ([]*entity.ArchiveVideo, error)
	Create(context.Context, *input.ArchiveVideoInput) error
	Update(context.Context, *entity.VideoId, *input.UpdateArchiveVideoInput) error
}
