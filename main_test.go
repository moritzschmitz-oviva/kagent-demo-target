package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Placeholder for the existing mainHandler in main.go
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// Placeholder for the new liveHandler to be implemented in main.go
func liveHandler(w http.ResponseWriter, r *http.Request) {
	// The actual implementation in main.go will have time.Sleep(1 * time.Second)
	// and w.WriteHeader(http.StatusOK)
	// This placeholder will make the tests fail because it doesn't do that yet.
	fmt.Fprintf(w, "") // Return an empty body for the placeholder
	w.WriteHeader(http.StatusTeapot) // Make it explicitly fail to ensure TDD
}

// This helper sets up a test HTTP server with the handlers
func setupTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler) // Existing handler
	mux.HandleFunc("/live", liveHandler) // New handler for /live
	return httptest.NewServer(mux)
}

func TestMainHandler(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	resp, err := server.Client().Get(server.URL + "/")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if string(body) != "Hello, World!" {
		t.Errorf("Expected body \"Hello, World!\"; got %q", body)
	}
}

func TestLiveEndpointReturns200AndHasDelay(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	start := time.Now()
	resp, err := server.Client().Get(server.URL + "/live")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	duration := time.Since(start)

	// This status code check will fail because our placeholder liveHandler returns StatusTeapot
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK for /live; got %v", resp.StatusCode)
	}

	// This delay check will fail because our placeholder liveHandler doesn't have a delay
	if duration < 1*time.Second || duration > 1500*time.Millisecond {
		t.Errorf("Expected /live endpoint to take approx 1 second; got %v", duration)
	}

	body, _ := io.ReadAll(resp.Body)
	if len(body) != 0 {
		t.Errorf("Expected empty body for /live; got %q", body)
	}
}

func TestLiveEndpointNoContent(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	resp, err := server.Client().Get(server.URL + "/live")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if len(body) != 0 {
		t.Errorf("Expected empty body for /live; got %q", body)
	}
}

func TestNonExistentEndpointReturns404(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	resp, err := server.Client().Get(server.URL + "/nonexistent")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status Not Found for /nonexistent; got %v", resp.StatusCode)
	}
}
