package validator

import "fizzbuzz/internal/fizzbuzz/domain/model"

type Validator interface {
	Validate(model.FizzBuzzDTO) error
}
