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
	t.Run("オプションなし", func(t *testing.T) {
		client := batch.NewHttpClient(http.MethodGet, nil)
		twitchApiClient := NewTwitchApiClient(client, cfg.NewConfig())

		listBroadcast, err := twitchApiClient.ListBroadcast(nil)
		a.NoError(err)
		a.NotNil(listBroadcast)
	})
	t.Run("オプションあり", func(t *testing.T) {
		client := batch.NewHttpClient(http.MethodGet, nil)
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

func Test_ListVideoByUserId(t *testing.T) {
	a := assert.New(t)

	// pekepokosanのユーザID
	searchUserId := "186620619"

	t.Run("オプションなし", func(t *testing.T) {
		client := batch.NewHttpClient(http.MethodGet, nil)
		twitchApiClient := NewTwitchApiClient(client, cfg.NewConfig())

		listBroadcast, err := twitchApiClient.ListVideoByUserId(searchUserId, nil)
		a.NoError(err)
		a.NotNil(listBroadcast)
	})
	t.Run("オプションあり", func(t *testing.T) {
		client := batch.NewHttpClient(http.MethodGet, nil)
		twitchApiClient := NewTwitchApiClient(client, cfg.NewConfig())

		options := []batch.RequestParam{
			{Key: "first", Value: "1"},
		}

		ListVideoByUserId, err := twitchApiClient.ListVideoByUserId(searchUserId, options)
		a.NoError(err)
		a.NotNil(ListVideoByUserId)
		a.Len(ListVideoByUserId.List, 1)

		actual := ListVideoByUserId.List[0]
		a.Equal("1473824992", actual.Id)
		a.Equal("39300467239", actual.StreamId)
		a.Equal("186620619", actual.UserId)
		a.Equal("ぺけぽこ", actual.UserName)
		a.NotNil(actual.Title)
		a.NotNil(actual.Viewable)
	})
}
