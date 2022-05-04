package live

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
)

type archiveVideoRepository struct {
	conn *postgresql.PostgreSql
}

func NewArchiveVideoRepository(conn *postgresql.PostgreSql) repository.ArchiveVideoRepositoryInterface {
	return &archiveVideoRepository{
		conn: conn,
	}
}

func (repository *archiveVideoRepository) Create(ctx context.Context, in *input.CreateArchiveVideoInput) error {
	err := repository.conn.Create(in)
	if err != nil {
		return err
	}

	return nil
}

func (repository *archiveVideoRepository) Get(ctx context.Context) (*entity.ArchiveVideo, error) {
	// TODO: 開発する
	return nil, nil
}

func (repository *archiveVideoRepository) List(ctx context.Context) ([]*entity.ArchiveVideo, error) {
	// TODO: 開発する
	return nil, nil
}
