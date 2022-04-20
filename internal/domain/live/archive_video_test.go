package live

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewArchiveVideo(t *testing.T) {
	a := assert.New(t)

	id := NewArchiveVideoId(1)

	archiveVideo := NewArchiveVideo(id)
	a.Equal(1, archiveVideo.Id().Int())
}
