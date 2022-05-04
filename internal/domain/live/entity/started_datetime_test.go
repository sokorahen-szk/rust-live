package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewStartedDatetime(t *testing.T) {
	a := assert.New(t)

	startedDatetime := NewStartedDatetime("2022-04-01T00:00:00Z")
	a.Equal("2022-04-01T00:00:00Z", startedDatetime.RFC3339())
	a.Equal("2022-04-01", startedDatetime.YYYYMMDD())
}

func Test_NewStartedDatetimeFromTime(t *testing.T) {
	a := assert.New(t)

	datetime := time.Date(2022, 5, 4, 23, 59, 59, 0, time.UTC)
	d := NewStartedDatetimeFromTime(&datetime)
	a.Equal("2022-05-04T23:59:59Z", d.RFC3339())
	a.Equal("2022-05-04", d.YYYYMMDD())
	a.Equal("2022-05-04 23:59:59", d.Timestamp())
	a.IsType(&time.Time{}, d.Time())
}
