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
		Addr:         ":8080",
		ReadTimeout:  TimeoutDuration,
		WriteTimeout: TimeoutDuration,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
