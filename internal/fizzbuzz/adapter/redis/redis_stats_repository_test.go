package redisrepo

import (
	"context"
	"fizzbuzz/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisStatsRepository(t *testing.T) {
	r := New(config.LocalHostRedisAddr, config.AlternateRedisTestDB)
	ctx := context.Background()
	type testCase struct {
		name        string
		prepare     func()
		expectKey   string
		expectCount int
	}

	tests := []testCase{
		{
			name: "single key incremented once",
			prepare: func() {
				r.IncrementKey(ctx, "test:key1")
			},
			expectKey:   "test:key1",
			expectCount: 1,
		},
		{
			name: "key incremented multiple times",
			prepare: func() {
				r.IncrementKey(ctx, "test:key2")
				r.IncrementKey(ctx, "test:key2")
				r.IncrementKey(ctx, "test:key2")
			},
			expectKey:   "test:key2",
			expectCount: 3,
		},
	}

	clearTestKeys(r, "test", ctx)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			key, count, err := r.GetMostUsedKey(ctx)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectKey, key)
			assert.Equal(t, tt.expectCount, count)
		})
	}
}

func clearTestKeys(r *RedisStatsRepository, prefix string, ctx context.Context) {
	keys, _ := r.client.Keys(ctx, prefix+"*").Result()
	if len(keys) > 0 {
		r.client.Del(ctx, keys...)
	}
}
