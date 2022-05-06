package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewPlatform(t *testing.T) {
	a := assert.New(t)

	platform := NewPlatform(PlatformTwitch)
	a.Equal(1, platform.Int())
	a.Equal("Twitch", platform.String())

	platform = NewPlatform(PlatformYoutube)
	a.Equal(2, platform.Int())
	a.Equal("Youtube", platform.String())
}

func Test_NewPlatformFromInt(t *testing.T) {
	a := assert.New(t)

	platform := NewPlatformFromInt(1)
	a.Equal(1, platform.Int())
	a.Equal("Twitch", platform.String())

	platform = NewPlatformFromInt(2)
	a.Equal(2, platform.Int())
	a.Equal("Youtube", platform.String())
}
