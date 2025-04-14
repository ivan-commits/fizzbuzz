package redisrepo

import (
	"context"
	"fizzbuzz/config"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisStatsRepository(t *testing.T) {
	r := New(config.LocalHostRedisAddr, config.AlternateRedisTestDB)

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
				r.IncrementKey("test:key1")
			},
			expectKey:   "test:key1",
			expectCount: 1,
		},
		{
			name: "key incremented multiple times",
			prepare: func() {
				r.IncrementKey("test:key2")
				r.IncrementKey("test:key2")
				r.IncrementKey("test:key2")
			},
			expectKey:   "test:key2",
			expectCount: 3,
		},
	}

	clearTestKeys(r, "test")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			key, count, err := r.GetMostUsedKey()
			assert.NoError(t, err)
			assert.Equal(t, tt.expectKey, key)
			assert.Equal(t, tt.expectCount, count)
		})
	}
}

func clearTestKeys(r *RedisStatsRepository, prefix string) {
	keys, _ := r.client.Keys(context.Background(), prefix+"*").Result()
	fmt.Println(keys)
	if len(keys) > 0 {
		r.client.Del(context.Background(), keys...)
	}
}
