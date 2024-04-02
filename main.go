package main

import (
	"fmt"
	"net/http"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.ListenAndServe(":8081", nil)
}
