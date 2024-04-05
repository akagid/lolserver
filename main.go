package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	ReadTimeout       = 5 * time.Second
	WriteTimeout      = 10 * time.Second
	ReadHeaderTimeout = 5 * time.Second
	IdleTimeout       = 60 * time.Second
)

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		ReadTimeout:       ReadTimeout,
		WriteTimeout:      WriteTimeout,
		ReadHeaderTimeout: ReadHeaderTimeout,
		IdleTimeout:       IdleTimeout,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
