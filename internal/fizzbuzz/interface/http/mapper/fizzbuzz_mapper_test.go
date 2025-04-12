package mapper

import (
	"fizzbuzz/internal/fizzbuzz/domain/model"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewFizzBuzzDTO(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected model.FizzBuzzDTO
	}{
		{
			name:  "valid parameters",
			query: "/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz",
			expected: model.FizzBuzzDTO{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "fizz",
				Str2:  "buzz",
			},
		},
		{
			name:  "missing parameters (zero values fallback)",
			query: "/fizzbuzz",
			expected: model.FizzBuzzDTO{
				Int1:  0,
				Int2:  0,
				Limit: 0,
				Str1:  "",
				Str2:  "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest("GET", tt.query, nil)
			c.Request = req

			dto := NewFizzBuzzDTO(c)
			assert.Equal(t, tt.expected, dto)
		})
	}
}
