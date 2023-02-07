package batch

import (
	"context"
)

type TwitchFetchLiveVideosUsecaseInterface interface {
	Handle(context.Context) error
}
