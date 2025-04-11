package port

import "fizzbuzz/internal/fizzbuzz/domain/model"

type Generator interface {
	Generate(dto model.FizzBuzzDTO) []string
}
