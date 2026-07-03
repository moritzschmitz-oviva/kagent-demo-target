package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is 200 OK.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body.
	expected := "Hello, World!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), expected)
	}
}

func TestHelloHandlerNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/not-found", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is 404 Not Found.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

// TestPingHandler tests the /ping endpoint
func TestPingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pingHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is 200 OK.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("pingHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body.
	expected := "{\"status\":\"ok\"}"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("pingHandler returned unexpected body: got %q want %q",
			rr.Body.String(), expected)
	}

	// Check Content-Type header
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("pingHandler returned wrong Content-Type: got %q want %q",
			contentType, "application/json")
	}
}

// TestPingHandlerMethodNotAllowed tests that /ping only allows GET requests
func TestPingHandlerMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("POST", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pingHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is 405 Method Not Allowed.
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("pingHandler with POST returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}
