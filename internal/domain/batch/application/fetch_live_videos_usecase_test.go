package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchLiveVideosUsecase_Handle(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()

	listLiveUsecase := NewInjectFetchLiveVideosUsecase(ctx)
	err := listLiveUsecase.Handle(ctx)
	a.NoError(err)
}
