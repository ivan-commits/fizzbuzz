package domain

import (
	"errors"
	"unicode/utf8"
)

func (r FizzBuzzRequest) IsValid() error {
	if r.Int1 <= 0 || r.Int2 <= 0 || r.Limit <= 0 {
		return errors.New("int1, int2, and limit must be > 0")
	}

	if utf8.RuneCountInString(r.Str1) == 0 || utf8.RuneCountInString(r.Str1) > 30 {
		return errors.New("str1 must be between 1 and 30 characters")
	}

	if utf8.RuneCountInString(r.Str2) == 0 || utf8.RuneCountInString(r.Str2) > 30 {
		return errors.New("str2 must be between 1 and 30 characters")
	}

	return nil
}
