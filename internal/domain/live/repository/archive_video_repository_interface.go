package repository

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
)

type ArchiveVideoRepositoryInterface interface {
	List(context.Context) ([]*entity.ArchiveVideo, error)
	Save(context.Context) error
}
