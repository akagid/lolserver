package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
)

func Add(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
