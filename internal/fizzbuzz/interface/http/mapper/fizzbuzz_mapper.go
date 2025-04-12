package mapper

import (
	model "fizzbuzz/internal/fizzbuzz/domain/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewFizzBuzzDTO(c *gin.Context) model.FizzBuzzDTO {
	return model.FizzBuzzDTO{
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
