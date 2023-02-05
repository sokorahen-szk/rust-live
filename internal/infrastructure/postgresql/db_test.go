package postgresql

import (
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/stretchr/testify/assert"
)

func Test_NewPostgreSQL(t *testing.T) {
	a := assert.New(t)
	postgresql := NewPostgreSQL(cfg.NewConfig())
	a.NotNil(postgresql)
	a.NoError(postgresql.Error)
}

func Test_NewPostgreSQL_Truncate(t *testing.T) {
	a := assert.New(t)
	postgresql := NewPostgreSQL(cfg.NewConfig())
	a.NotNil(postgresql)
	a.NoError(postgresql.Error)

	err := postgresql.Truncate([]string{"archive_videos"})
	a.NoError(err)
}
