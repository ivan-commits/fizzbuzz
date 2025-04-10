package stats

import (
	"fizzbuzz/pkg/model"
	"sync"
)

var (
	mu       sync.Mutex
	counter  = map[model.FizzBuzzRequest]int{}
	topHit   model.FizzBuzzRequest
	maxCount int
)

func SaveRequest(req model.FizzBuzzRequest) {
	mu.Lock()
	defer mu.Unlock()
	counter[req]++
	if counter[req] > maxCount {
		maxCount = counter[req]
		topHit = req
	}
}

func GetMostUsed() (model.FizzBuzzRequest, int) {
	mu.Lock()
	defer mu.Unlock()
	return topHit, maxCount
}
