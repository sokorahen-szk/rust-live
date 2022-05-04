package twitch

import (
	"net/http"
	"testing"

	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch"
	"github.com/stretchr/testify/assert"

	cfg "github.com/sokorahen-szk/rust-live/config"
)

func Test_ListBroadcast(t *testing.T) {
	a := assert.New(t)
	url := "https://api.twitch.tv/helix/streams"

	t.Run("オプションなし", func(t *testing.T) {
		client := batch.NewHttpClient(http.MethodGet, url, nil)
		twitchApiClient := NewTwitchApiClient(client, cfg.NewConfig())

		listBroadcast, err := twitchApiClient.ListBroadcast(nil)
		a.NoError(err)
		a.NotNil(listBroadcast)
	})
	t.Run("オプションあり", func(t *testing.T) {
		client := batch.NewHttpClient(http.MethodGet, url, nil)
		twitchApiClient := NewTwitchApiClient(client, cfg.NewConfig())

		options := []batch.RequestParam{
			{Key: "language", Value: "ja"},
			{Key: "game_id", Value: RustGameId},
			{Key: "type", Value: "live"},
			{Key: "first", Value: "100"},
		}

		listBroadcast, err := twitchApiClient.ListBroadcast(options)
		a.NoError(err)
		a.NotNil(listBroadcast)
	})
}
