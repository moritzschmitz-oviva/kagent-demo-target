package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

func TestLiveEndpointReturns200AndHasDelay(t *testing.T) {
	req, err := http.NewRequest("GET", "/live", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	start := time.Now()
	handler.ServeHTTP(rr, req)
	duration := time.Since(start)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// The description specifies a 1-second delay. Allowing for some minor variance.
	if duration < 1*time.Second {
		t.Errorf("handler returned too quickly: got %v, expected at least %v",
			duration, 1*time.Second)
	}
}

func TestLiveEndpointNoContent(t *testing.T) {
	req, err := http.NewRequest("GET", "/live", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	expected := "" // Assuming no content for a liveness probe
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUnknownEndpointReturns404(t *testing.T) {
	req, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
