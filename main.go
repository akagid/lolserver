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

func newServer() *http.Server {
	return &http.Server{
		Addr: ":8080",
	}
}

func setupRoutes() {
	http.HandleFunc("/", helloHandler)
}

func main() {
	setupRoutes()

	err := newServer().ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
