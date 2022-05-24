package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewElapsedTimes(t *testing.T) {
	a := assert.New(t)

	t.Run("1時間の場合", func(t *testing.T) {
		elaspedTimes := NewElapsedTimes(3600)
		a.Equal(3600, elaspedTimes.Int())
		a.Equal(int32(3600), elaspedTimes.Int32())
		a.Equal(1, elaspedTimes.Hours())
		a.Equal(0, elaspedTimes.Minutes())
		a.Equal(0, elaspedTimes.Seconds())
		a.Equal("01h00m00s", elaspedTimes.Text())
	})

	t.Run("59分59秒の場合", func(t *testing.T) {
		elaspedTimes := NewElapsedTimes(3599)
		a.Equal(3599, elaspedTimes.Int())
		a.Equal(int32(3599), elaspedTimes.Int32())
		a.Equal(0, elaspedTimes.Hours())
		a.Equal(59, elaspedTimes.Minutes())
		a.Equal(59, elaspedTimes.Seconds())
		a.Equal("59m59s", elaspedTimes.Text())
	})

	t.Run("59秒の場合", func(t *testing.T) {
		elaspedTimes := NewElapsedTimes(59)
		a.Equal(59, elaspedTimes.Int())
		a.Equal(int32(59), elaspedTimes.Int32())
		a.Equal(0, elaspedTimes.Hours())
		a.Equal(0, elaspedTimes.Minutes())
		a.Equal(59, elaspedTimes.Seconds())
		a.Equal("59s", elaspedTimes.Text())
	})
}
