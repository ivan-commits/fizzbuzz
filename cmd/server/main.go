package main

import (
	"fizzbuzz/cmd/router"
	"fizzbuzz/config"
	redisrepo "fizzbuzz/internal/fizzbuzz/adapter/redis"
	"fizzbuzz/internal/fizzbuzz/usecase"
)

func main() {

	Generator := usecase.DefaultGenerator{}
	StatsRepository := redisrepo.New(config.DefaultRedisAddr, config.DefaultRedisDB)
	Validator := usecase.DefaultValidator{}

	r := router.NewRouter(Generator, StatsRepository, Validator)
	r.Run(":8000")
}
