package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	readTimeout       = 5 * time.Second
	writeTimeout      = 10 * time.Second
	readHeaderTimeout = 5 * time.Second
	idleTimeout       = 60 * time.Second
)

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
