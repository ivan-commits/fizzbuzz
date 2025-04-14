package handler

import (
	"context"
	"errors"
	"fizzbuzz/internal/fizzbuzz/domain/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGenerator struct{ mock.Mock }

func (m *MockGenerator) Generate(dto model.FizzBuzzDTO) []string {
	args := m.Called(dto)
	return args.Get(0).([]string)
}

type MockStatsRepository struct{ mock.Mock }

func (m *MockStatsRepository) IncrementKey(ctx context.Context, key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *MockStatsRepository) GetMostUsedKey(ctx context.Context) (string, int, error) {
	args := m.Called()
	return args.String(0), args.Int(1), args.Error(2)
}

type MockValidator struct{ mock.Mock }

func (m *MockValidator) Validate(dto model.FizzBuzzDTO) error {
	args := m.Called(dto)
	return args.Error(0)
}

func TestHandler_HandleFizzbuzz(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		query        string
		mockSetup    func(*MockGenerator, *MockStatsRepository, *MockValidator)
		expectedCode int
	}{
		{
			name:  "success",
			query: "/fizzbuzz?int1=3&int2=5&limit=5&str1=fizz&str2=buzz",
			mockSetup: func(g *MockGenerator, s *MockStatsRepository, v *MockValidator) {
				dto := model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 5, Str1: "fizz", Str2: "buzz"}
				v.On("Validate", dto).Return(nil)
				g.On("Generate", dto).Return([]string{"1", "2", "fizz", "4", "buzz"})
				s.On("IncrementKey", mock.Anything).Return(nil)
			},
			expectedCode: http.StatusOK,
		},
		{
			name:  "invalid input",
			query: "/fizzbuzz?int1=0&int2=0&limit=0&str1=&str2=",
			mockSetup: func(g *MockGenerator, s *MockStatsRepository, v *MockValidator) {
				v.On("Validate", mock.Anything).Return(errors.New("invalid"))
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := new(MockGenerator)
			s := new(MockStatsRepository)
			v := new(MockValidator)
			if tt.mockSetup != nil {
				tt.mockSetup(g, s, v)
			}

			h := NewHandler(g, s, v)
			r := gin.Default()
			r.GET("/fizzbuzz", h.HandleFizzbuzz)

			req := httptest.NewRequest("GET", tt.query, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestHandler_HandleStats(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := new(MockStatsRepository)
	s.On("GetMostUsedKey").Return("key", 10, nil)

	h := NewHandler(nil, s, nil)
	r := gin.Default()
	r.GET("/stats", h.HandleStats)

	req := httptest.NewRequest("GET", "/stats", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
