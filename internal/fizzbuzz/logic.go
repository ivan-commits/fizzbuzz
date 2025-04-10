package fizzbuzz

import (
	"fizzbuzz/pkg/model"
	"strconv"
)

func Generate(req model.FizzBuzzRequest) []string {
	res := make([]string, 0, req.Limit)
	for i := 1; i <= req.Limit; i++ {
		switch {
		case i%req.Int1 == 0 && i%req.Int2 == 0:
			res = append(res, req.Str1+req.Str2)
		case i%req.Int1 == 0:
			res = append(res, req.Str1)
		case i%req.Int2 == 0:
			res = append(res, req.Str2)
		default:
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
