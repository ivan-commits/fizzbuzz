package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"fizzbuzz/cmd/router"
	"fizzbuzz/internal/fizzbuzz/usecase"

	"github.com/stretchr/testify/assert"
)

type MockStatsRepository struct{}

func (ms MockStatsRepository) IncrementKey(ctx context.Context, key string) error {
	return nil
}

func (ms MockStatsRepository) GetMostUsedKey(ctx context.Context) (string, int, error) {
	return "test", 2, nil
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
		assert.JSONEq(t, `{"key":"test","count":2}`, w.Body.String())
		assert.True(t, w.Code == http.StatusOK)
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
