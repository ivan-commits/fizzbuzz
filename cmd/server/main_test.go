package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"fizzbuzz/cmd/router"
	"fizzbuzz/internal/fizzbuzz/usecase"

	"github.com/stretchr/testify/assert"
)

type MockStatsRepository struct{}

func (ms MockStatsRepository) IncrementKey(key string) error {
	return nil
}

func (ms MockStatsRepository) GetMostUsedKey() (string, int, error) {
	return "key", 2, nil
}

func TestMainRouter_Endpoints(t *testing.T) {

	Generator := usecase.DefaultGenerator{}
	StatsRepository := MockStatsRepository{}
	Validator := usecase.DefaultValidator{}

	r := router.NewRouter(Generator, StatsRepository, Validator)

	t.Run("GET /stats", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/stats", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.True(t, w.Code == http.StatusOK || w.Code == http.StatusBadRequest)
	})

	t.Run("GET /fizzbuzz", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GET /fizzbuzz", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/fizzbuzz", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
