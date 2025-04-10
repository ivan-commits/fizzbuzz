package handler

import (
	"encoding/json"
	"fizzbuzz/internal/fizzbuzz"
	"fizzbuzz/internal/stats"
	"fizzbuzz/pkg/model"
	"net/http"
	"strconv"
)

func FizzBuzz(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	req := model.FizzBuzzRequest{
		Int1:  parseInt(q.Get("int1")),
		Int2:  parseInt(q.Get("int2")),
		Limit: parseInt(q.Get("limit")),
		Str1:  q.Get("str1"),
		Str2:  q.Get("str2"),
	}
	result := fizzbuzz.Generate(req)
	stats.SaveRequest(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func Stats(w http.ResponseWriter, r *http.Request) {
	top, hits := stats.GetMostUsed()
	resp := map[string]interface{}{
		"request": top,
		"hits":    hits,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
