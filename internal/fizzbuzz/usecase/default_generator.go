package usecase

import (
	domain "fizzbuzz/internal/fizzbuzz/domain/model"
	"strconv"
)

type DefaultGenerator struct{}

func (dg DefaultGenerator) Generate(dto domain.FizzBuzzDTO) []string {
	res := make([]string, 0, dto.Limit)
	for i := 1; i <= dto.Limit; i++ {
		switch {
		case i%dto.Int1 == 0 && i%dto.Int2 == 0:
			res = append(res, dto.Str1+dto.Str2)
		case i%dto.Int1 == 0:
			res = append(res, dto.Str1)
		case i%dto.Int2 == 0:
			res = append(res, dto.Str2)
		default:
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
