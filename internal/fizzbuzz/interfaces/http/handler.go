package http

import (
	"fizzbuzz/internal/fizzbuzz/domain/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Generator       port.Generator
	StatsRepository port.StatsRepository
	Validator       port.Validator
}

func NewHandler(Generator port.Generator, StatsRepository port.StatsRepository, Validator port.Validator) *Handler {
	return &Handler{Generator: Generator, StatsRepository: StatsRepository, Validator: Validator}
}

func (h *Handler) HandleFizzbuzz(c *gin.Context) {

	dto := NewFizzBuzzDTO(c)
	err := h.Validator.Validate(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.Generator.Generate(dto)

	res := NewFizzBuzzResponse(dto, result)

	key := CacheKey(dto)

	err = h.StatsRepository.IncrementKey(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) HandleStats(c *gin.Context) {
	key, count, err := h.StatsRepository.GetMostUsedKey()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sr := NewStatsResponse(key, count)

	c.JSON(http.StatusOK, sr)
}
