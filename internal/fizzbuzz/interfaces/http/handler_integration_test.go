package http

import (
	"fizzbuzz/internal/fizzbuzz/application"
	model "fizzbuzz/internal/fizzbuzz/domain/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandlerGET_Valid_WithMock(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup mock
	mockGen := new(MockGenerator)
	expected := model.FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 10,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	mockGen.On("Generate", expected).Return([]string{"1", "2", "fizz", "4", "buzz"})

	// Setup handler
	handler := NewHandler(mockGen)
	router := gin.Default()
	router.GET("/fizzbuzz", handler.GET)

	// Test request
	req := httptest.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=10&str1=fizz&str2=buzz", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	mockGen.AssertCalled(t, "Generate", expected)
	assert.JSONEq(t, `{"response": ["1", "2", "fizz", "4", "buzz"]}`, w.Body.String())
}

func TestHandlerGET_Invalid(t *testing.T) {
	gin.SetMode(gin.TestMode)

	generator := application.DefaultGenerator{}
	handler := NewHandler(&generator)
	router := gin.Default()
	router.GET("/fizzbuzz", handler.GET)

	req := httptest.NewRequest("GET", "/fizzbuzz", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestHandlerGET_LimitNegative(t *testing.T) {
	gin.SetMode(gin.TestMode)

	generator := application.DefaultGenerator{}
	handler := NewHandler(&generator)

	router := gin.Default()
	router.GET("/fizzbuzz", handler.GET)

	req := httptest.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=-1&str1=fizz&str2=buzz", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "limit")
}

func TestHandlerGET_Str1TooLong(t *testing.T) {
	gin.SetMode(gin.TestMode)

	generator := application.DefaultGenerator{}
	handler := NewHandler(&generator)

	router := gin.Default()
	router.GET("/fizzbuzz", handler.GET)

	longStr := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" // 53 char
	req := httptest.NewRequest("GET", "/fizzbuzz?int1=1&int2=2&limit=10&str1="+longStr+"&str2=buzz", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "str1")
}

func TestHandlerGET_LimitOne(t *testing.T) {
	gin.SetMode(gin.TestMode)

	generator := application.DefaultGenerator{}
	handler := NewHandler(&generator)

	router := gin.Default()
	router.GET("/fizzbuzz", handler.GET)

	req := httptest.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=1&str1=fizz&str2=buzz", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"response": ["1"]}`, w.Body.String())
}
