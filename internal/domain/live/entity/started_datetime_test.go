package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewStartedDatetime(t *testing.T) {
	a := assert.New(t)

	startedDatetime := NewStartedDatetime("2022-04-01T00:00:00Z")
	a.Equal("2022-04-01T00:00:00Z", startedDatetime.RFC3339())
	a.Equal("2022-04-01", startedDatetime.YYYYMMDD())
}
