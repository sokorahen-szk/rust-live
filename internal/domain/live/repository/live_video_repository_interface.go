package repository

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
)

type CacheLiveVideoListKey string

const (
	TwitchLiveVideoKey  CacheLiveVideoListKey = "twitch_live_video_list"
	YoutubeLiveVideoKey CacheLiveVideoListKey = "youtube_live_video_lists"
)

type LiveVideoRepositoryInterface interface {
	Create(context.Context, []*entity.LiveVideo, CacheLiveVideoListKey) error
	List(context.Context, *list.ListLiveVideoInput, CacheLiveVideoListKey) ([]*entity.LiveVideo, error)
	Count(context.Context, CacheLiveVideoListKey) (int, error)
	Analytics(context.Context, CacheLiveVideoListKey) (*live.AnalyticsOutput, error)
}
