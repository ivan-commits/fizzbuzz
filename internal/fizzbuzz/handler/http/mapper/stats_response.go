package mapper

import (
	model "fizzbuzz/internal/fizzbuzz/domain/model"
	"fmt"
)

type StatsReponse struct {
	Key   string `json:"key"`
	Count int    `json:"count"`
}

func GenerateCacheKey(r model.FizzBuzzDTO) string {
	return fmt.Sprintf("int1=%d&int2=%d&limit=%d&str1=%s&str2=%s",
		r.Int1, r.Int2, r.Limit, r.Str1, r.Str2)
}

func NewStatsResponse(key string, count int) StatsReponse {
	return StatsReponse{Key: key, Count: count}
}
