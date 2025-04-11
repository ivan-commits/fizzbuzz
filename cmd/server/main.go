package main

import (
	"fizzbuzz/cmd/router"
)

func main() {
	r := router.NewRouter()
	r.Run(":8000")
}
