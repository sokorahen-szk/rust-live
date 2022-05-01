package live

import (
	"context"
	"encoding/json"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
)

type liveVideoRepository struct {
	conn *redis.Redis
}

const liveVideoListKey = "live_video_list_key"

func NewLiveVideoRepository(conn *redis.Redis) repository.LiveVideoRepositoryInterface {
	return &liveVideoRepository{
		conn: conn,
	}
}

func (repository *liveVideoRepository) List(ctx context.Context, listInput *list.ListLiveVideosInput) ([]*entity.LiveVideo, error) {
	data, err := repository.conn.Get(ctx, liveVideoListKey)
	if err != nil {
		return nil, err
	}

	var liveVideos []*entity.LiveVideo
	err = json.Unmarshal([]byte(data), &liveVideos)
	if err != nil {
		return nil, err
	}

	return liveVideos, nil
}

func (repository *liveVideoRepository) Create(ctx context.Context, liveVideos []*entity.LiveVideo) error {
	serializeData, err := json.Marshal(&liveVideos)
	if err != nil {
		return err
	}

	setData := &redis.RedisSetData{
		Key:   liveVideoListKey,
		Value: serializeData,
		Ttl:   nil,
	}
	err = repository.conn.Set(ctx, setData)
	if err != nil {
		return err
	}

	return nil
}
