package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	TimeoutDuration = 10 * time.Second
)

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		ReadTimeout:       TimeoutDuration,
		WriteTimeout:      TimeoutDuration,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
