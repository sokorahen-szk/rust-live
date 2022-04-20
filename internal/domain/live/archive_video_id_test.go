package live

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewArchiveVideoId(t *testing.T) {
	a := assert.New(t)

	archiveVideoId := NewArchiveVideoId(1)
	a.Equal(1, archiveVideoId.Int())
}
