package batch

import (
	"context"
)

type UpdateLiveVideosUsecaseInterface interface {
	Handle(context.Context) error
}
