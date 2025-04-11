package http

import (
	domain "fizzbuzz/internal/fizzbuzz/domain/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Generator domain.Generator
}

func NewHandler(Generator domain.Generator) *Handler {
	return &Handler{Generator: Generator}
}

func (h *Handler) GET(c *gin.Context) {

	req := ParseRequest(c)
	err := req.IsValid()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.Generator.Generate(req)

	res := NewFizzBuzzResponse(req, result)

	c.JSON(http.StatusOK, res)
}
