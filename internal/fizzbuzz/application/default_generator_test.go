package application

import (
	domain "fizzbuzz/internal/fizzbuzz/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultGenerator_Generate(t *testing.T) {
	generator := DefaultGenerator{}

	req := domain.FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 15,
		Str1:  "fizz",
		Str2:  "buzz",
	}

	expected := []string{
		"1", "2", "fizz", "4", "buzz",
		"fizz", "7", "8", "fizz", "buzz",
		"11", "fizz", "13", "14", "fizzbuzz",
	}

	result := generator.Generate(req)
	assert.Equal(t, expected, result)
}

func TestDefaultGenerator_LimitZero(t *testing.T) {
	gen := DefaultGenerator{}
	req := domain.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 0, Str1: "fizz", Str2: "buzz"}

	result := gen.Generate(req)
	assert.Empty(t, result)
}

func TestDefaultGenerator_EqualDivisors(t *testing.T) {
	gen := DefaultGenerator{}
	req := domain.FizzBuzzRequest{Int1: 2, Int2: 2, Limit: 4, Str1: "fizz", Str2: "buzz"}

	result := gen.Generate(req)
	assert.Equal(t, []string{"1", "fizzbuzz", "3", "fizzbuzz"}, result)
}

func TestDefaultGenerator_SameStrings(t *testing.T) {
	gen := DefaultGenerator{}
	req := domain.FizzBuzzRequest{Int1: 2, Int2: 3, Limit: 6, Str1: "go", Str2: "go"}

	result := gen.Generate(req)
	assert.Equal(t, []string{"1", "go", "go", "go", "5", "gogo"}, result)
}

func TestDefaultGenerator_LimitOne(t *testing.T) {
	gen := DefaultGenerator{}
	req := domain.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 1, Str1: "fizz", Str2: "buzz"}

	result := gen.Generate(req)
	assert.Equal(t, []string{"1"}, result)
}
