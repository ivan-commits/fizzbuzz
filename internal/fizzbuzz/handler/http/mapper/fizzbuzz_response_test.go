package mapper

import (
	"fizzbuzz/internal/fizzbuzz/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFizzBuzzResponse(t *testing.T) {
	tests := []struct {
		name     string
		inputDTO model.FizzBuzzDTO
		inputRes []string
		expected FizzBuzzResponse
	}{
		{
			name:     "standard fizzbuzz response",
			inputDTO: model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 5, Str1: "fizz", Str2: "buzz"},
			inputRes: []string{"1", "2", "fizz", "4", "buzz"},
			expected: FizzBuzzResponse{Response: []string{"1", "2", "fizz", "4", "buzz"}},
		},
		{
			name:     "empty result",
			inputDTO: model.FizzBuzzDTO{},
			inputRes: []string{},
			expected: FizzBuzzResponse{Response: []string{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewFizzBuzzResponse(tt.inputDTO, tt.inputRes)
			assert.Equal(t, tt.expected, result)
		})
	}
}
