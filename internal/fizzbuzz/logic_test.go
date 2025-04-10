package fizzbuzz

import (
	"fizzbuzz/pkg/model"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	req := model.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"}
	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}

	result := Generate(req)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
