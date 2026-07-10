package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cenkalti/backoff/v4"
)

// simulateJiraAPIRequest simulates a call to the Jira API that might fail.
// It will succeed on the 3rd attempt.
func simulateJiraAPIRequest(attempt int) error {
	if attempt < 3 {
		return fmt.Errorf("simulated Jira API error on attempt %d", attempt)
	}
	return nil
}

// CallJiraAPIWithRetry calls a simulated Jira API with exponential backoff.
func CallJiraAPIWithRetry() error {
	operation := func() error {
		log.Println("Attempting to call Jira API...")
		return simulateJiraAPIRequest(operationAttempt)
	}

	var operationAttempt int
	notify := func(err error, d time.Duration) {
		operationAttempt++
		log.Printf("Jira API call failed: %v. Retrying in %s...", err, d)
	}

	exponentialBackOff := backoff.NewExponentialBackOff()
	exponentialBackOff.MaxElapsedTime = 30 * time.Second // Max 30 seconds for all retries

	err := backoff.RetryNotify(operation, exponentialBackOff, notify)
	if err != nil {
		log.Printf("Failed to call Jira API after multiple retries: %v", err)
		return err
	}

	log.Println("Successfully called Jira API after retries.")
	return nil
}
