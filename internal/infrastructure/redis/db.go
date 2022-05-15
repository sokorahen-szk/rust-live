package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	cfg "github.com/sokorahen-szk/rust-live/config"
)

type Redis struct {
	ctx        context.Context
	db         *redis.Client
	defaultTtl time.Duration
}

type RedisSetData struct {
	Key   string
	Value interface{}
	Ttl   *time.Duration
}

type RedisCacheEmptyError struct {
	error
}

func (e *RedisCacheEmptyError) Is(err error) bool {
	return true
}

func NewRedis(ctx context.Context, c *cfg.Config) *Redis {
	addr := fmt.Sprintf(
		"%s:%d",
		c.Redis.Host,
		c.Redis.Port,
	)
	db := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: c.Redis.Password,
		DB:       c.Redis.DbNumber,
	})

	ttl := time.Hour * time.Duration(c.Redis.DefaultTtlHour)
	return &Redis{
		ctx:        ctx,
		db:         db,
		defaultTtl: ttl,
	}
}

func (redis *Redis) Set(ctx context.Context, setData *RedisSetData) error {
	ttl := redis.defaultTtl
	if setData.Ttl != nil {
		ttl = *setData.Ttl
	}

	statusCmd := redis.db.Set(ctx, setData.Key, setData.Value, ttl)
	return statusCmd.Err()
}

func (redis *Redis) Get(ctx context.Context, key string) (string, error) {
	val, err := redis.db.Get(ctx, key).Result()

	if err != nil && redis.isCacheEmpty(err) {
		return "", RedisCacheEmptyError{err}
	}

	return val, err
}

func (redis *Redis) Truncate() error {
	statusCmd := redis.db.FlushAll(redis.ctx)
	return statusCmd.Err()
}

func (redis *Redis) isCacheEmpty(err error) bool {
	if err.Error() == "redis: nil" {
		return true
	}
	return false
}
