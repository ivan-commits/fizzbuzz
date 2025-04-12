package usecase

import (
	"fizzbuzz/internal/fizzbuzz/domain/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultValidator_Validate(t *testing.T) {
	type args struct {
		dto model.FizzBuzzDTO
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "valid input",
			args:    args{dto: model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"}},
			wantErr: false,
		},
		{
			name:    "zero int1",
			args:    args{dto: model.FizzBuzzDTO{Int1: 0, Int2: 5, Limit: 10, Str1: "fizz", Str2: "buzz"}},
			wantErr: true,
		},
		{
			name:    "zero int2",
			args:    args{dto: model.FizzBuzzDTO{Int1: 3, Int2: 0, Limit: 10, Str1: "fizz", Str2: "buzz"}},
			wantErr: true,
		},
		{
			name:    "zero limit",
			args:    args{dto: model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 0, Str1: "fizz", Str2: "buzz"}},
			wantErr: true,
		},
		{
			name:    "empty str1",
			args:    args{dto: model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 10, Str1: "", Str2: "buzz"}},
			wantErr: true,
		},
		{
			name:    "empty str2",
			args:    args{dto: model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 10, Str1: "fizz", Str2: ""}},
			wantErr: true,
		},
		{
			name:    "str1 too long",
			args:    args{dto: model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 10, Str1: strings.Repeat("a", 31), Str2: "buzz"}},
			wantErr: true,
		},
		{
			name:    "str2 too long",
			args:    args{dto: model.FizzBuzzDTO{Int1: 3, Int2: 5, Limit: 10, Str1: "fizz", Str2: strings.Repeat("b", 31)}},
			wantErr: true,
		},
	}

	validator := DefaultValidator{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(tt.args.dto)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
