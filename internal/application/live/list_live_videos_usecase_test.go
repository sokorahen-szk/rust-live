package application_live

import (
	"context"
	"testing"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/redis"
	redisLive "github.com/sokorahen-szk/rust-live/internal/infrastructure/redis/live"
	"github.com/sokorahen-szk/rust-live/internal/usecase/live/list"
	mockEntity "github.com/sokorahen-szk/rust-live/tests/domain/live/entity"
	"github.com/stretchr/testify/assert"
)

func Test_ListLiveVideosUsecase_Handle(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	platforms := []*entity.Platform{}
	sortKey := entity.NewLiveVideoSortKey(0)

	config := cfg.NewConfig()
	redis := redis.NewRedis(ctx, config)
	liveVideoRepository := redisLive.NewLiveVideoRepository(redis)

	usecase := NewListLiveVideosUsecase(liveVideoRepository)

	t.Run("liveVideoが0件の時、空を返すこと", func(t *testing.T) {
		redis.Truncate()

		input := list.NewListLiveVideoInput("", platforms, sortKey, 1, 10)

		expectedPagination := &pb.Pagination{
			Limit:      10,
			Page:       1,
			Prev:       1,
			Next:       1,
			TotalPage:  0,
			TotalCount: 0,
		}

		actual, err := usecase.Handle(ctx, input)
		a.NotNil(actual)
		a.NoError(err)
		a.Len(actual.LiveVideos, 0)
		a.Equal(expectedPagination, actual.Pagination)
	})
	t.Run("liveVideoが存在する時、LiveVideo配列を返すこと", func(t *testing.T) {
		redis.Truncate()

		input := list.NewListLiveVideoInput("", platforms, sortKey, 1, 10)

		liveVideos := []*entity.LiveVideo{
			mockEntity.NewMockLiveVideo(1),
			mockEntity.NewMockLiveVideo(2),
			mockEntity.NewMockLiveVideo(3),
		}

		expectedPagination := &pb.Pagination{
			Limit:      10,
			Page:       1,
			Prev:       1,
			Next:       1,
			TotalPage:  1,
			TotalCount: 3,
		}

		err := liveVideoRepository.Create(ctx, liveVideos, repository.TwitchLiveVideoKey)
		a.NoError(err)

		actual, err := usecase.Handle(ctx, input)
		a.NotNil(actual)
		a.NoError(err)
		a.Len(actual.LiveVideos, 3)
		a.Equal(expectedPagination, actual.Pagination)
	})
}
