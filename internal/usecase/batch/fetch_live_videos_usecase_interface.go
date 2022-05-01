package batch

import (
	"context"
)

type FetchLiveVideosUsecaseInterface interface {
	Handle(context.Context) error
}
