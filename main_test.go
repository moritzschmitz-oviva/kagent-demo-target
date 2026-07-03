package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time" // Added time package
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
	// This will fail compilation as liveHandler doesn't exist yet, which is expected.
	handler := http.HandlerFunc(liveHandler)

	start := time.Now()
	handler.ServeHTTP(rr, req)
	elapsed := time.Since(start)

	// Check the status code is 200 OK.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the delay is approximately 1 second.
	// Allowing for some minor overhead, so between 1 and 1.5 seconds.
	if elapsed < 1*time.Second || elapsed > 1500*time.Millisecond {
		t.Errorf("handler returned too quickly or too slowly: got %v, expected around 1 second", elapsed)
	}
}

func TestLiveEndpointNoContent(t *testing.T) {
	req, err := http.NewRequest("GET", "/live", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	// This will fail compilation as liveHandler doesn't exist yet, which is expected.
	handler := http.HandlerFunc(liveHandler)

	handler.ServeHTTP(rr, req)

	// Check the response body is empty.
	if rr.Body.String() != "" {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), "")
	}
}
