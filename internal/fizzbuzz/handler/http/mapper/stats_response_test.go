package mapper

import (
	"fizzbuzz/internal/fizzbuzz/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCacheKey(t *testing.T) {
	tests := []struct {
		name     string
		input    model.FizzBuzzDTO
		expected string
	}{
		{
			name:     "standard input",
			input:    model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"},
			expected: "int1=3&int2=5&limit=15&str1=fizz&str2=buzz",
		},
		{
			name:     "empty input",
			input:    model.FizzBuzzDTO{},
			expected: "int1=0&int2=0&limit=0&str1=&str2=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateCacheKey(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewStatsResponse(t *testing.T) {
	key := "int1=3&int2=5&limit=15&str1=fizz&str2=buzz"
	count := 42

	expected := StatsReponse{
		Key:   key,
		Count: count,
	}

	result := NewStatsResponse(key, count)
	assert.Equal(t, expected, result)
}
