package usecase

import (
	"fizzbuzz/internal/fizzbuzz/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultGenerator_Generate(t *testing.T) {
	type args struct {
		dto model.FizzBuzzDTO
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "standard fizzbuzz",
			args: args{dto: model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"}},
			want: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
		},
		{
			name: "same string",
			args: args{dto: model.FizzBuzzDTO{Int1: 2, Int2: 3, Limit: 6, Str1: "go", Str2: "go"}},
			want: []string{"1", "go", "go", "go", "5", "gogo"},
		},
		{
			name: "limit 0",
			args: args{dto: model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 0, Str1: "fizz", Str2: "buzz"}},
			want: []string{},
		},
		{
			name: "equal divisors",
			args: args{dto: model.FizzBuzzDTO{Int1: 2, Int2: 2, Limit: 4, Str1: "fizz", Str2: "buzz"}},
			want: []string{"1", "fizzbuzz", "3", "fizzbuzz"},
		},
		{
			name: "limit 1",
			args: args{dto: model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 1, Str1: "fizz", Str2: "buzz"}},
			want: []string{"1"},
		},
	}

	gen := DefaultGenerator{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gen.Generate(tt.args.dto)
			assert.Equal(t, tt.want, got)
		})
	}
}
