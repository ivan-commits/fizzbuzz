package redisrepo

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisStatsRepository struct {
	client *redis.Client
}

var ctx = context.Background()

func New(addr string, db int) *RedisStatsRepository {
	return &RedisStatsRepository{
		client: redis.NewClient(&redis.Options{Addr: addr, DB: db}),
	}
}

func (r *RedisStatsRepository) IncrementKey(key string) error {
	return r.client.Incr(ctx, key).Err()
}

func (r *RedisStatsRepository) GetMostUsedKey() (string, int, error) {
	keys, err := r.client.Keys(ctx, "*").Result()

	if err != nil {
		return "", 0, err
	}

	var maxKey string
	var maxVal int

	for _, key := range keys {
		val, err := r.client.Get(ctx, key).Int()
		if err == nil && val > maxVal {
			maxVal = val
			maxKey = key
		}
	}

	return maxKey, maxVal, nil
}
