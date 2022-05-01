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
	key   string
	value interface{}
	ttl   *time.Duration
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

func (redis *Redis) Set(setData *RedisSetData) error {
	ttl := redis.defaultTtl
	if setData.ttl != nil {
		ttl = *setData.ttl
	}

	statusCmd := redis.db.Set(redis.ctx, setData.key, setData.value, ttl)
	return statusCmd.Err()
}

func (redis *Redis) Get(key string) (string, error) {
	val, err := redis.db.Get(redis.ctx, key).Result()
	return val, err
}

func (redis *Redis) Truncate() error {
	statusCmd := redis.db.FlushAll(redis.ctx)
	return statusCmd.Err()
}
