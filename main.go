package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func liveHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/live", liveHandler) // Register the new liveHandler
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
