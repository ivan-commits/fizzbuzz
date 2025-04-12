package handler

import (
	"fizzbuzz/internal/fizzbuzz/domain/contract"
	"fizzbuzz/internal/fizzbuzz/interface/http/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Generator       contract.Generator
	StatsRepository contract.StatsRepository
	Validator       contract.Validator
}

func NewHandler(Generator contract.Generator, StatsRepository contract.StatsRepository, Validator contract.Validator) *Handler {
	return &Handler{Generator: Generator, StatsRepository: StatsRepository, Validator: Validator}
}

func (h *Handler) HandleFizzbuzz(c *gin.Context) {

	dto := mapper.NewFizzBuzzDTO(c)
	err := h.Validator.Validate(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.Generator.Generate(dto)

	res := mapper.NewFizzBuzzResponse(dto, result)

	key := mapper.GenerateCacheKey(dto)

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

	sr := mapper.NewStatsResponse(key, count)

	c.JSON(http.StatusOK, sr)
}
