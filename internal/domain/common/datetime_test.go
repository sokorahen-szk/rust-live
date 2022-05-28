package common

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewDatetime(t *testing.T) {
	a := assert.New(t)

	t.Run("RFC3339をNewDatetimeに渡す場合", func(t *testing.T) {
		d := NewDatetime("2022-04-01T00:00:00Z")
		a.Equal("2022-04-01T00:00:00Z", d.RFC3339())
		a.Equal("2022-04-01", d.YYYYMMDD())
		a.Equal("2022-04-01 00:00:00", d.Timestamp())
		a.IsType(&time.Time{}, d.Time())
	})
	t.Run("timestampをNewDatetimeに渡す場合", func(t *testing.T) {
		d := NewDatetime("2022-04-01 00:00:00")
		a.Equal("2022-04-01T00:00:00Z", d.RFC3339())
		a.Equal("2022-04-01", d.YYYYMMDD())
		a.Equal("2022-04-01 00:00:00", d.Timestamp())
		a.IsType(&time.Time{}, d.Time())
	})
}

func Test_NewDatetimeFromTime(t *testing.T) {
	a := assert.New(t)

	datetime := time.Date(2022, 5, 4, 23, 59, 59, 0, time.UTC)
	d := NewDatetimeFromTime(&datetime)
	a.Equal("2022-05-04T23:59:59Z", d.RFC3339())
	a.Equal("2022-05-04", d.YYYYMMDD())
	a.Equal("2022-05-04 23:59:59", d.Timestamp())
	a.IsType(&time.Time{}, d.Time())
}

func Test_DiffSeconds(t *testing.T) {
	a := assert.New(t)

	t.Run("d > d2の時、正の値を返す", func(t *testing.T) {
		d := NewDatetime("2022-02-02T01:01:10Z")
		d2 := NewDatetime("2022-02-02T00:00:00Z")
		actual := d.DiffSeconds(d2.Time())
		a.Equal(3670, actual)
	})
	t.Run("d < d2の時、負の値を返す", func(t *testing.T) {
		d := NewDatetime("2022-02-02T00:00:00Z")
		d2 := NewDatetime("2022-02-02T01:01:10Z")
		actual := d.DiffSeconds(d2.Time())
		a.Equal(-3670, actual)
	})
	t.Run("d == d2の時、0を返す", func(t *testing.T) {
		d := NewDatetime("2022-02-02T00:00:00Z")
		d2 := NewDatetime("2022-02-02T00:00:00Z")
		actual := d.DiffSeconds(d2.Time())
		a.Equal(0, actual)
	})
}

func Test_Lt(t *testing.T) {
	a := assert.New(t)

	t.Run("x > b のとき、falseになること", func(t *testing.T) {
		x := NewDatetime("2022-02-02T00:00:01Z")
		y := NewDatetime("2022-02-02T00:00:00Z")

		a.False(x.Lt(y.Time()))
	})
	t.Run("x < b のとき、trueになること", func(t *testing.T) {
		x := NewDatetime("2022-02-02T00:00:00Z")
		y := NewDatetime("2022-02-02T00:00:01Z")

		a.True(x.Lt(y.Time()))
	})
	t.Run("x = b のとき、falseになること", func(t *testing.T) {
		x := NewDatetime("2022-02-02T00:00:00Z")
		y := NewDatetime("2022-02-02T00:00:00Z")

		a.False(x.Lt(y.Time()))
	})
}

func Test_Gt(t *testing.T) {
	a := assert.New(t)

	t.Run("x > b のとき、trueになること", func(t *testing.T) {
		x := NewDatetime("2022-02-02T00:00:01Z")
		y := NewDatetime("2022-02-02T00:00:00Z")

		a.True(x.Gt(y.Time()))
	})
	t.Run("x < b のとき、falseになること", func(t *testing.T) {
		x := NewDatetime("2022-02-02T00:00:00Z")
		y := NewDatetime("2022-02-02T00:00:01Z")

		a.False(x.Gt(y.Time()))
	})
	t.Run("x = b のとき、falseになること", func(t *testing.T) {
		x := NewDatetime("2022-02-02T00:00:00Z")
		y := NewDatetime("2022-02-02T00:00:00Z")

		a.False(x.Lt(y.Time()))
	})
}
