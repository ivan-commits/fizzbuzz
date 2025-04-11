package domain

import (
	domain "fizzbuzz/internal/fizzbuzz/domain/model"
)

// Generator est une abstraction de la logique FizzBuzz.
//
// Elle encapsule la règle métier suivante :
// Pour chaque entier de 1 à Limit inclus :
//   - si divisible par Int1 et Int2 → retourner Str1 + Str2
//   - sinon si divisible par Int1   → retourner Str1
//   - sinon si divisible par Int2   → retourner Str2
//   - sinon                          → retourner le nombre lui-même
//
// Cette interface permet de substituer différentes implémentations
// (optimisée, distribuée, mockée...) sans modifier le handler.
type Generator interface {
	Generate(req domain.FizzBuzzRequest) []string
}
