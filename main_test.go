package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello, World!"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expected)
	}
}

func TestNewServer(t *testing.T) {
	t.Parallel()

	server := newServer()

	if server.Addr != ":8080" {
		t.Errorf("server.Addr = %s; want :8080", server.Addr)
	}

	if server.ReadTimeout != readTimeout {
		t.Errorf("server.ReadTimeout = %v; want %v", server.ReadTimeout, readTimeout)
	}

	if server.WriteTimeout != writeTimeout {
		t.Errorf("server.WriteTimeout = %v; want %v", server.WriteTimeout, writeTimeout)
	}

	if server.ReadHeaderTimeout != readHeaderTimeout {
		t.Errorf("server.ReadHeaderTimeout = %v; want %v", server.ReadHeaderTimeout, readHeaderTimeout)
	}

	if server.IdleTimeout != idleTimeout {
		t.Errorf("server.IdleTimeout = %v; want %v", server.IdleTimeout, idleTimeout)
	}
}

func TestSetupRoutes(t *testing.T) {
	t.Parallel()

	setupRoutes()

	// Get the default ServeMux
	mux := http.DefaultServeMux

	// Check if "/" route is registered
	_, pattern := mux.Handler(&http.Request{URL: &url.URL{Path: "/"}})
	if pattern != "/" {
		t.Errorf("Expected route to be '/', got %s", pattern)
	}
}
