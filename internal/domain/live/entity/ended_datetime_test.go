package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewEndedDatetime(t *testing.T) {
	a := assert.New(t)

	endedDatetime := NewEndedDatetime("2022-04-01T00:00:00Z")
	a.Equal("2022-04-01T00:00:00Z", endedDatetime.RFC3339())
	a.Equal("2022-04-01", endedDatetime.YYYYMMDD())
}
