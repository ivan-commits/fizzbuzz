package http

import (
	model "fizzbuzz/internal/fizzbuzz/domain/model"
)

type FizzBuzzResponse struct {
	Response []string `json:"response"`
}

func NewFizzBuzzResponse(req model.FizzBuzzDTO, res []string) FizzBuzzResponse {
	return FizzBuzzResponse{Response: res}
}
