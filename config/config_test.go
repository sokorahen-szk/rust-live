package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewConfig(t *testing.T) {
	a := assert.New(t)
	c := NewConfig()

	t.Run(".envファイルがLoadできること", func(t *testing.T) {
		a.Equal("local", c.Env)
		a.Equal(9000, c.Port)
	})
}
