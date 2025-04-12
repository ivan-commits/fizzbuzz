package main

import (
	"fizzbuzz/cmd/router"
	redisrepo "fizzbuzz/internal/fizzbuzz/adapter/redis"
	"fizzbuzz/internal/fizzbuzz/usecase"
)

func main() {

	Generator := usecase.DefaultGenerator{}
	StatsRepository := redisrepo.New("redis:6379", 1)
	Validator := usecase.DefaultValidator{}

	r := router.NewRouter(Generator, StatsRepository, Validator)
	r.Run(":8000")
}
