package http

import (
	domain "fizzbuzz/internal/fizzbuzz/domain/model"

	"github.com/stretchr/testify/mock"
)

type MockGenerator struct {
	mock.Mock
}

func (m *MockGenerator) Generate(req domain.FizzBuzzRequest) []string {
	args := m.Called(req)
	return args.Get(0).([]string)
}
