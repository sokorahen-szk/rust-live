package youtube

import (
	"strings"

	cfg "github.com/sokorahen-szk/rust-live/config"
	httpClient "github.com/sokorahen-szk/rust-live/pkg/http"
)

const (
	// カテゴリーID
	// https://gist.github.com/dgp/1b24bf2961521bd75d6c
	CategoryGameId         string = "20"
	ListBroadcastUrl       string = "https://www.googleapis.com/youtube/v3/search"
	ListVideoByVideoIdsUrl string = "https://www.googleapis.com/youtube/v3/videos"
)

type YouTubeApiClientInterface interface {
	ListBroadcast() (*ListBroadcastResponse, error)
	ListVideoByVideoIds([]string) (*ListVideoByVideoIdsResponse, error)
}

type ListBroadcastResponse struct {
	List []ListBroadcast `json:"items"`
}

type ListBroadcast struct {
	Id      ListBroadcastId     `json:"id"`
	Snippet ListBroadcastSippet `json:"snippet"`
}

type ListBroadcastId struct {
	VideoId string `json:"videoId"` // streamId
}

type ListBroadcastSippet struct {
	PublishedAt  string                 `json:"publishedAt"` // startedAt
	Title        string                 `json:"title"`
	ThumbnailUrl ListBroadcastThumbnail `json:"thumbnails"`
	ChannelId    string                 `json:"channelId"`    // userId
	ChannelTitle string                 `json:"channelTitle"` // userName
}
type ListBroadcastThumbnail struct {
	High ListBroadcastThumbnailHigh `json:"high"`
}
type ListBroadcastThumbnailHigh struct {
	Url string `json:"url"`
}

type ListVideoByVideoIdsResponse struct {
	List []ListVideoByVideoIds `json:"items"`
}

type ListVideoByVideoIds struct {
	VideoId string                    `json:"id"`
	Detail  ListVideoByVideoIdsDetail `json:"liveStreamingDetails"`
}
type ListVideoByVideoIdsDetail struct {
	ViewerCount string `json:"concurrentViewers"`
}

type YoutubeApiClient struct {
	httpClient httpClient.HttpClientInterface

	auth string
}

func NewYoutubeApiClient(httpClient httpClient.HttpClientInterface, config *cfg.Config) YouTubeApiClientInterface {
	apiClient := &YoutubeApiClient{
		httpClient: httpClient,
	}

	apiClient.auth = config.Batch.ApiYoutubeSecretKey
	apiClient.setAuth()

	return apiClient
}

func (api *YoutubeApiClient) ListBroadcast() (*ListBroadcastResponse, error) {
	api.httpClient.AddParams([]httpClient.RequestParam{
		{Key: "regionCode", Value: "jp"},
		{Key: "q", Value: "rust"},
		{Key: "part", Value: "snippet"},
		{Key: "eventType", Value: "live"},
		{Key: "type", Value: "video"},
		{Key: "maxResults", Value: "50"},
		{Key: "videoCategoryId", Value: CategoryGameId},
	})

	httpClientGetResponse, err := api.httpClient.Get(ListBroadcastUrl, &ListBroadcastResponse{})
	if err != nil {
		return nil, err
	}

	listBroadcastResponse := httpClientGetResponse.Data.(*ListBroadcastResponse)
	return listBroadcastResponse, nil
}

func (api *YoutubeApiClient) ListVideoByVideoIds(videoIds []string) (*ListVideoByVideoIdsResponse, error) {
	params := []httpClient.RequestParam{
		{Key: "part", Value: "liveStreamingDetails"},
		{Key: "id", Value: strings.Join(videoIds, ",")},
	}

	api.httpClient.AddParams(params)

	httpClientGetResponse, err := api.httpClient.Get(ListVideoByVideoIdsUrl, &ListVideoByVideoIdsResponse{})
	if err != nil {
		return nil, err
	}

	listVideoByVideoIdsResponse := httpClientGetResponse.Data.(*ListVideoByVideoIdsResponse)
	return listVideoByVideoIdsResponse, nil
}

func (api *YoutubeApiClient) setAuth() {
	api.httpClient.AddParams([]httpClient.RequestParam{
		{Key: "key", Value: api.auth},
	})
}
