package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFizzBuzzHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz", nil)
	rec := httptest.NewRecorder()

	FizzBuzz(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "fizzbuzz") {
		t.Errorf("Response missing expected fizzbuzz content: %s", body)
	}
}

func TestStatsHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/stats", nil)
	rec := httptest.NewRecorder()

	Stats(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), `"hits":`) {
		t.Errorf("Expected stats in response, got: %s", body)
	}
}
