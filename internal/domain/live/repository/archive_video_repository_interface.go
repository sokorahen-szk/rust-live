package repository

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
)

type ArchiveVideoRepositoryInterface interface {
	Get(context.Context) (*entity.ArchiveVideo, error)
	List(context.Context) ([]*entity.ArchiveVideo, error)
	Create(context.Context, *input.CreateArchiveVideoInput) error
}
