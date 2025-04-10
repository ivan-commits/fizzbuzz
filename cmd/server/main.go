package main

import (
	"fizzbuzz/internal/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/fizzbuzz", handler.FizzBuzz)
	mux.HandleFunc("/stats", handler.Stats)

	log.Println("Listening on :8000...")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
