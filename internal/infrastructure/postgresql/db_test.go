package postgresql

import (
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/stretchr/testify/assert"
)

func Test_NewPostgreSQL(t *testing.T) {
	a := assert.New(t)
	postgresql := NewPostgreSQL(cfg.NewConfig())
	a.NotNil(postgresql.db)
	a.NoError(postgresql.db.Error)
}
