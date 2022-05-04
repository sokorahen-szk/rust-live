package live

import (
	"context"
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/domain/common"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
	"github.com/stretchr/testify/assert"
)

func Test_ArchiveVideoRepository_Create(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	postgresql := postgresql.NewPostgreSQL(cfg.NewConfig())

	repository := NewArchiveVideoRepository(postgresql)

	time := common.NewDatetime("2022-02-02T14:00:00Z")

	in := &input.CreateArchiveVideoInput{
		BroadcastId:     "39300467239",
		Title:           "title",
		Url:             "https://example.com/test",
		Stremer:         "テスター",
		ThumbnailImage:  "https://example.com/test.jpg",
		StartedDatetime: time.Time(),
	}

	err := repository.Create(ctx, in)
	a.NoError(err)
}
