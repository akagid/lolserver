package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
