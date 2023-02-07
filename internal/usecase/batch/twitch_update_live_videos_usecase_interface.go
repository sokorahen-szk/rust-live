package batch

import (
	"context"
)

type TwitchUpdateLiveVideosUsecaseInterface interface {
	Handle(context.Context) error
}
