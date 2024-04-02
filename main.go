package main

import (
	"fmt"
	"net/http"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/piing", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "poong")

		e := 1
	})

	http.ListenAndServe(":8081", nil)
}
