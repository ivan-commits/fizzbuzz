package application

import (
	"errors"
	domain "fizzbuzz/internal/fizzbuzz/domain/model"
	"unicode/utf8"
)

type DefaultValidator struct{}

func (dv DefaultValidator) Validate(dto domain.FizzBuzzDTO) error {
	if dto.Int1 <= 0 || dto.Int2 <= 0 || dto.Limit <= 0 {
		return errors.New("int1, int2, and limit must be > 0")
	}

	if utf8.RuneCountInString(dto.Str1) == 0 || utf8.RuneCountInString(dto.Str1) > 30 {
		return errors.New("str1 must be between 1 and 30 characters")
	}

	if utf8.RuneCountInString(dto.Str2) == 0 || utf8.RuneCountInString(dto.Str2) > 30 {
		return errors.New("str2 must be between 1 and 30 characters")
	}

	return nil
}
