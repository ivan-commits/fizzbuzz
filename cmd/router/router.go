package router

import (
	"fizzbuzz/internal/fizzbuzz/domain/contract"
	"fizzbuzz/internal/fizzbuzz/handler/http/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(g contract.Generator, s contract.StatsRepository, v contract.Validator) *gin.Engine {
	r := gin.Default()

	Handler := handler.NewHandler(g, s, v)

	r.GET("/fizzbuzz", Handler.HandleFizzbuzz)
	r.GET("/stats", Handler.HandleStats)
	return r
}
