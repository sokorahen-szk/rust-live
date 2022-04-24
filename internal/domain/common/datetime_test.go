package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDatetime(t *testing.T) {
	a := assert.New(t)

	d := NewDatetime("2022-04-01T00:00:00Z")
	a.Equal("2022-04-01T00:00:00Z", d.RFC3339())
	a.Equal("2022-04-01", d.YYYYMMDD())
}
