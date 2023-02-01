package twitch

import (
	"fmt"

	cfg "github.com/sokorahen-szk/rust-live/config"
	httpClient "github.com/sokorahen-szk/rust-live/pkg/http"
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
	StreamId     string `json:"id"`
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

type TwitchApiClientInterface interface {
	ListBroadcast([]httpClient.RequestParam) (*ListBroadcastResponse, error)
	ListVideoByUserId(string, []httpClient.RequestParam) (*ListVideoByUserIdResponse, error)
}
type TwitchApiClient struct {
	httpClient httpClient.HttpClientInterface

	clientId string
	auth     string
}

func NewTwitchApiClient(httpClient httpClient.HttpClientInterface, config *cfg.Config) TwitchApiClientInterface {
	apiClient := &TwitchApiClient{
		httpClient: httpClient,
	}

	apiClient.clientId = config.Batch.ApiTwitchClientId
	apiClient.auth = config.Batch.ApiTwtichSecretKey
	apiClient.setAuth()

	return apiClient
}

func (api *TwitchApiClient) ListBroadcast(params []httpClient.RequestParam) (*ListBroadcastResponse, error) {
	api.httpClient.AddParams(params)
	httpClientGetResponse, err := api.httpClient.Get(ListBroadcastUrl, &ListBroadcastResponse{})
	if err != nil {
		return nil, err
	}

	ListBroadcastResponse := httpClientGetResponse.Data.(*ListBroadcastResponse)
	return ListBroadcastResponse, nil
}

func (api *TwitchApiClient) ListVideoByUserId(userId string, params []httpClient.RequestParam) (*ListVideoByUserIdResponse, error) {
	params = append(params, []httpClient.RequestParam{
		{Key: "user_id", Value: userId},
	}...)
	api.httpClient.AddParams(params)
	httpClientGetResponse, err := api.httpClient.Get(GetVideoByUserUrl, &ListVideoByUserIdResponse{})
	if err != nil {
		return nil, err
	}

	listVideoByUserIdResponse := httpClientGetResponse.Data.(*ListVideoByUserIdResponse)
	return listVideoByUserIdResponse, nil
}

func (api *TwitchApiClient) setAuth() {
	bearer := fmt.Sprintf("Bearer %s", api.auth)
	api.httpClient.AddHeaders([]httpClient.RequestHeader{
		{Key: "Client-Id", Value: api.clientId},
		{Key: "Authorization", Value: bearer},
	})
}
