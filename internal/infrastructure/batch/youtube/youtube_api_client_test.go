package youtube

import (
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	httpClient "github.com/sokorahen-szk/rust-live/pkg/http"
	"github.com/stretchr/testify/assert"
)

func TestListBroadcast(t *testing.T) {
	a := assert.New(t)

	client := httpClient.NewHttpClient(nil)
	twitchApiClient := NewYoutubeApiClient(client, cfg.NewConfig())

	listBroadcast, err := twitchApiClient.ListBroadcast()
	a.NoError(err)
	a.NotNil(listBroadcast)
}

func TestListVideoByVideoIds(t *testing.T) {
	a := assert.New(t)

	client := httpClient.NewHttpClient(nil)
	twitchApiClient := NewYoutubeApiClient(client, cfg.NewConfig())

	videoIds := []string{"47DLMopLG9I", "T5THTPQw1SE"}
	listVideoByVideoIds, err := twitchApiClient.ListVideoByVideoIds(videoIds)
	a.NoError(err)
	a.NotNil(listVideoByVideoIds)
}
