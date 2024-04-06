// main package
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	readHeaderTimeout = 5 * time.Second
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

func run() error {
	setupRoutes()

	err := newServer().ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
