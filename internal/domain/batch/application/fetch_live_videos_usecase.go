package application

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/batch/twitch"
	usecaseBatch "github.com/sokorahen-szk/rust-live/internal/usecase/batch"
)

type fetchLiveVideosUsecase struct {
	liveVideoRepository repository.LiveVideoRepositoryInterface
	twitchApiClient     twitch.TwitchApiClientInterface
}

func NewFetchLiveVideosUsecase(
	liveVideoRepository repository.LiveVideoRepositoryInterface,
	twitchApiClient twitch.TwitchApiClientInterface,
) usecaseBatch.FetchLiveVideosUsecaseInterface {
	return fetchLiveVideosUsecase{
		liveVideoRepository: liveVideoRepository,
		twitchApiClient:     twitchApiClient,
	}
}

func (usecase fetchLiveVideosUsecase) Handle(ctx context.Context) error {
	/*
		options := []batch.RequestParam{
			{Key: "language", Value: "ja"},
			{Key: "game_id", Value: twitch.RustGameId},
			{Key: "type", Value: "live"},
			{Key: "first", Value: "100"},
		}

		listBroadcastResponse, err := usecase.twitchApiClient.ListBroadcast(options)
		if err != nil {
			return err
		}
	*/
	return nil
}
