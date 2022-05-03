package twitch

import (
	"fmt"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch"
)

type ListBroadcastResponse struct {
	data *ListBroadcastResponseData
}

type ListBroadcastResponseData struct {
	id string `json:"id`
}

type twitchApiClient struct {
	httpClient *batch.HttpClient

	clientId string
	auth     string
}

func NewTwitchApiClient(httpClient *batch.HttpClient, config *cfg.Config) batch.ApiClient {
	apiClient := &twitchApiClient{
		httpClient: httpClient,
	}

	apiClient.clientId = config.Batch.ApiTwitchClientId
	apiClient.auth = config.Batch.ApiTwtichSecretKey
	apiClient.SetAuth()

	return apiClient
}

func (api *twitchApiClient) SetAuth() {
	bearer := fmt.Sprintf("Bearer %s", api.auth)
	api.httpClient.AddHeaders([]batch.RequestHeader{
		{Key: "Client-Id", Value: api.clientId},
		{Key: "Authorization", Value: bearer},
	})
}

func (api *twitchApiClient) ListBroadcast() {
	// TODO:
}
