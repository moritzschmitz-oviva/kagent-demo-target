package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	// Call the Jira API with retry logic
	if err := CallJiraAPIWithRetry(); err != nil {
		log.Fatalf("Jira API client failed: %v", err)
	}

	http.HandleFunc("/", helloHandler)
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
