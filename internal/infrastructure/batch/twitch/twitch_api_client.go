package twitch

import (
	"fmt"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch"
)

const (
	RustGameId        string = "263490"
	ListBroadcastUrl  string = "https://api.twitch.tv/helix/streams"
	GetVideoByUserUrl string = "https://api.twitch.tv/helix/videos"
)

type ListBroadcastResponse struct {
	List []ListBroadcast `json:"data"`
}

type ListBroadcast struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	UserLogin    string `json:"user_login"`
	UserName     string `json:"user_name"`
	Title        string `json:"title"`
	ViewerCount  int    `json:"viewer_count"`
	StartedAt    string `json:"started_at"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

type ListVideoByUserIdResponse struct {
	List []ListVideoByUserId `json:"data"`
}

type ListVideoByUserId struct {
	Id       string `json:"id"`
	StreamId string `json:"stream_id"`
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Title    string `json:"title"`
	Viewable string `json:"viewable"`
}

type twitchApiClient struct {
	httpClient *batch.HttpClient

	clientId string
	auth     string
}

func NewTwitchApiClient(httpClient *batch.HttpClient, config *cfg.Config) *twitchApiClient {
	apiClient := &twitchApiClient{
		httpClient: httpClient,
	}

	apiClient.clientId = config.Batch.ApiTwitchClientId
	apiClient.auth = config.Batch.ApiTwtichSecretKey
	apiClient.setAuth()

	return apiClient
}

func (api *twitchApiClient) ListBroadcast(params []batch.RequestParam) (*ListBroadcastResponse, error) {
	api.httpClient.AddParams(params)
	httpClientGetResponse, err := api.httpClient.Get(ListBroadcastUrl, &ListBroadcastResponse{})
	if err != nil {
		return nil, err
	}

	api.httpClient.DeleteParams(params)

	ListBroadcastResponse := httpClientGetResponse.Data.(*ListBroadcastResponse)
	return ListBroadcastResponse, nil
}

func (api *twitchApiClient) ListVideoByUserId(userId string, params []batch.RequestParam) (*ListVideoByUserIdResponse, error) {
	params = append(params, []batch.RequestParam{
		{Key: "user_id", Value: userId},
	}...)
	api.httpClient.AddParams(params)
	httpClientGetResponse, err := api.httpClient.Get(GetVideoByUserUrl, &ListVideoByUserIdResponse{})
	if err != nil {
		return nil, err
	}

	api.httpClient.DeleteParams(params)

	listVideoByUserIdResponse := httpClientGetResponse.Data.(*ListVideoByUserIdResponse)
	return listVideoByUserIdResponse, nil
}

func (api *twitchApiClient) setAuth() {
	bearer := fmt.Sprintf("Bearer %s", api.auth)
	api.httpClient.AddHeaders([]batch.RequestHeader{
		{Key: "Client-Id", Value: api.clientId},
		{Key: "Authorization", Value: bearer},
	})
}
