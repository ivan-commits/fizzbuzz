package stats

import (
	"fizzbuzz/pkg/model"
	"testing"
)

func TestSaveAndGetMostUsed(t *testing.T) {
	req1 := model.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"}
	req2 := model.FizzBuzzRequest{Int1: 2, Int2: 4, Limit: 10, Str1: "foo", Str2: "bar"}

	// Reset global store (important en tests)
	counter = map[model.FizzBuzzRequest]int{}
	maxCount = 0
	topHit = model.FizzBuzzRequest{}

	SaveRequest(req1)
	SaveRequest(req2)
	SaveRequest(req2)

	top, hits := GetMostUsed()

	if top != req2 {
		t.Errorf("Expected top request to be %v, got %v", req2, top)
	}
	if hits != 2 {
		t.Errorf("Expected 2 hits for top request, got %d", hits)
	}
}
