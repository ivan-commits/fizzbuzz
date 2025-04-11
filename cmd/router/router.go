package router

import (
	"fizzbuzz/internal/fizzbuzz/application"
	redisrepo "fizzbuzz/internal/fizzbuzz/infrastructure/redis"
	httphandler "fizzbuzz/internal/fizzbuzz/interfaces/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	Generator := application.DefaultGenerator{}
	StatsRepository := redisrepo.New("redis:6379")
	Validator := application.DefaultValidator{}

	Handler := httphandler.NewHandler(Generator, StatsRepository, Validator)

	r.GET("/fizzbuzz", Handler.HandleFizzbuzz)
	r.GET("/stats", Handler.HandleStats)
	return r
}
