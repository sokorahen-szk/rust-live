package live

import (
	"context"
	"encoding/json"
	"errors"
	"sort"
	"strings"

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

func (repository *liveVideoRepository) List(ctx context.Context, listInput *list.ListLiveVideoInput) ([]*entity.LiveVideo, error) {
	var liveVideos []*entity.LiveVideo

	data, err := repository.conn.Get(ctx, liveVideoListKey)
	if err != nil {
		if errors.Is(&redis.RedisCacheEmptyError{}, err) {
			return liveVideos, nil
		}

		return nil, err
	}

	err = json.Unmarshal([]byte(data), &liveVideos)
	if err != nil {
		return nil, err
	}

	return repository.listFilter(liveVideos, listInput), nil
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

func (repository *liveVideoRepository) listFilter(liveVideos []*entity.LiveVideo, listInput *list.ListLiveVideoInput) []*entity.LiveVideo {
	filtered := make([]*entity.LiveVideo, 0)

	for _, liveVideo := range liveVideos {
		if !strings.Contains(liveVideo.GetTitle().String(), listInput.SearchKeywords()) &&
			!strings.Contains(liveVideo.GetStremer().String(), listInput.SearchKeywords()) {
			continue
		}

		if !repository.isTargetPlatform(liveVideo, listInput.Platforms()) {
			continue
		}

		filtered = append(filtered, liveVideo)
	}

	sorted := repository.sort(filtered, listInput.SortKey())
	return repository.paginate(sorted, listInput.Page(), listInput.Limit())
}

func (repository *liveVideoRepository) isTargetPlatform(liveVideo *entity.LiveVideo,
	platforms []*entity.Platform) bool {
	if platforms == nil || len(platforms) < 1 {
		return true
	}

	for _, platform := range platforms {
		if *platform == *liveVideo.GetPlatform() {
			return true
		}
	}

	return false
}

func (repository *liveVideoRepository) sort(liveVideos []*entity.LiveVideo, sortKey *entity.LiveVideoSortKey) []*entity.LiveVideo {
	if sortKey == nil {
		return liveVideos
	}

	switch *sortKey {
	case entity.LiveVideoViewerAsc:
		sort.SliceStable(liveVideos, func(i, j int) bool {
			return liveVideos[i].Viewer.Int() < liveVideos[j].Viewer.Int()
		})
	case entity.LiveVideoViewerDesc:
		sort.SliceStable(liveVideos, func(i, j int) bool {
			return liveVideos[i].Viewer.Int() > liveVideos[j].Viewer.Int()
		})
	case entity.LiveVideoStartedDatetimeAsc:
		sort.SliceStable(liveVideos, func(i, j int) bool {
			return liveVideos[i].StartedDatetime.RFC3339() < liveVideos[j].StartedDatetime.RFC3339()
		})
	case entity.LiveVideoStartedDatetimeDesc:
		sort.SliceStable(liveVideos, func(i, j int) bool {
			return liveVideos[i].StartedDatetime.RFC3339() > liveVideos[j].StartedDatetime.RFC3339()
		})
	}

	return liveVideos
}

func (repository *liveVideoRepository) paginate(listVideos []*entity.LiveVideo, page int, limit int) []*entity.LiveVideo {
	if page == 0 || limit == 0 {
		return listVideos
	}

	// 全データがリミットより小さいときは、そのまま返す.
	if len(listVideos) < limit {
		return listVideos
	}

	first := (page - 1) * limit
	end := page * limit

	if len(listVideos) < end {
		return listVideos[first:]
	}

	return listVideos[first:end]
}
