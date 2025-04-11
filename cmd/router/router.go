package router

import (
	"fizzbuzz/internal/fizzbuzz/application"
	httphandler "fizzbuzz/internal/fizzbuzz/interfaces/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	generator := application.DefaultGenerator{}

	handler := httphandler.NewHandler(&generator)

	r.GET("/fizzbuzz", handler.GET)

	return r
}
