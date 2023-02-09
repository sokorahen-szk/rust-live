package batch

import (
	"context"
)

type YoutubeFetchLiveVideosUsecaseInterface interface {
	Handle(context.Context) error
}
