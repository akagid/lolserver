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
		Addr:                         ":8080",
		Handler:                      http.DefaultServeMux,
		ReadTimeout:                  readTimeout,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 writeTimeout,
		IdleTimeout:                  0,
		MaxHeaderBytes:               http.DefaultMaxHeaderBytes,
		TLSConfig:                    nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
		DisableGeneralOptionsHandler: false,
		TLSNextProto:                 nil,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
