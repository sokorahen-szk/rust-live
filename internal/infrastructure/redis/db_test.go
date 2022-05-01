package redis

import (
	"context"
	"testing"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"github.com/stretchr/testify/assert"
)

func Test_NewRedis(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	redis := NewRedis(ctx, cfg.NewConfig())
	a.NotNil(redis.db)
}

func Test_Set(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	redis := NewRedis(ctx, cfg.NewConfig())

	data := &RedisSetData{"key", "value", nil}
	err := redis.Set(ctx, data)
	a.NoError(err)
}

func Test_Get(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	redis := NewRedis(ctx, cfg.NewConfig())

	expected := "value"
	searchKey := "key"
	actual, err := redis.Get(ctx, searchKey)
	a.NoError(err)
	a.Equal(expected, actual)
}

func Test_Truncate(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	redis := NewRedis(ctx, cfg.NewConfig())
	a.NoError(redis.Truncate())
}
