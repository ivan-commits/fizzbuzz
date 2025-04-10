package model

import (
	"encoding/json"
	"testing"
)

func TestFizzBuzzRequest_JSON(t *testing.T) {
	req := FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 15,
		Str1:  "fizz",
		Str2:  "buzz",
	}

	// Marshal en JSON
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Erreur lors de la sérialisation JSON : %v", err)
	}

	// Unmarshal dans une nouvelle instance
	var out FizzBuzzRequest
	err = json.Unmarshal(data, &out)
	if err != nil {
		t.Fatalf("Erreur lors du parsing JSON : %v", err)
	}

	if req != out {
		t.Errorf("Objet désérialisé différent de l'original. Got %+v, expected %+v", out, req)
	}
}

func TestFizzBuzzRequest_MapKey(t *testing.T) {
	m := make(map[FizzBuzzRequest]int)

	key := FizzBuzzRequest{Int1: 1, Int2: 2, Limit: 3, Str1: "a", Str2: "b"}
	m[key]++

	if m[key] != 1 {
		t.Errorf("FizzBuzzRequest ne fonctionne pas comme clé de map")
	}
}
