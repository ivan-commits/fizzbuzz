package http

import (
	model "fizzbuzz/internal/fizzbuzz/domain/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseRequest(c *gin.Context) model.FizzBuzzRequest {
	return model.FizzBuzzRequest{
		Int1:  atoi(c.Query("int1")),
		Int2:  atoi(c.Query("int2")),
		Limit: atoi(c.Query("limit")),
		Str1:  c.Query("str1"),
		Str2:  c.Query("str2"),
	}
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
