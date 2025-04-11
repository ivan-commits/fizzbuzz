package domain

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzRequest_IsValid_Valid(t *testing.T) {
	req := FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 15,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	err := req.IsValid()
	assert.Nil(t, err)
}

func TestFizzBuzzRequest_IsValid_Int1Invalid(t *testing.T) {
	req := FizzBuzzRequest{
		Int1:  0,
		Int2:  5,
		Limit: 10,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	err := req.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "int1")
}

func TestFizzBuzzRequest_IsValid_Int2Invalid(t *testing.T) {
	req := FizzBuzzRequest{
		Int1:  3,
		Int2:  0,
		Limit: 10,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	err := req.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "int2")
}

func TestFizzBuzzRequest_IsValid_LimitInvalid(t *testing.T) {
	req := FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 0,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	err := req.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "limit")
}

func TestFizzBuzzRequest_IsValid_Str1Empty(t *testing.T) {
	req := FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 10,
		Str1:  "",
		Str2:  "buzz",
	}
	err := req.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "str1")
}

func TestFizzBuzzRequest_IsValid_Str2Empty(t *testing.T) {
	req := FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 10,
		Str1:  "fizz",
		Str2:  "",
	}
	err := req.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "str2")
}

func TestFizzBuzzRequest_IsValid_Str1TooLong(t *testing.T) {
	longStr := string(make([]rune, 31)) // 31 caractères
	req := FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 10,
		Str1:  longStr,
		Str2:  "buzz",
	}
	assert.Greater(t, utf8.RuneCountInString(req.Str1), 30)
	err := req.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "str1")
}

func TestFizzBuzzRequest_IsValid_Str2TooLong(t *testing.T) {
	longStr := string(make([]rune, 31))
	req := FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 10,
		Str1:  "fizz",
		Str2:  longStr,
	}
	assert.Greater(t, utf8.RuneCountInString(req.Str2), 30)
	err := req.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "str2")
}
