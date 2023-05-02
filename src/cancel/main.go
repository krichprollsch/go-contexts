package main

import (
	"fmt"
	"time"
)

// slow simulates a slow operation, like an HTTP request.
func slow(cancel <-chan struct{}) {
	select {
	// Simulate a long operation, like an HTTP request.
	case <-time.After(2 * time.Second):
		// nothing to do here
	case <-cancel:
		// operation is canceled, return early.
		return
	}

	fmt.Println(" world!")
}

func main() {
	// Run a function synchronously.
	fmt.Print("Hello")

	// done is a channel used to synchronize the goroutine and the main program.
	done := make(chan struct{})

	// cancel is a channel used to send a cancelation message from the main
	// program to the go routine.
	cancel := make(chan struct{})

	// Run a function asynchronously, ie. in a goroutine.
	go func() {
		slow(cancel)
		// Once the function stops, we send a message in the channel.
		done <- struct{}{}
	}()

	// Cancel the goroutine immediately.
	cancel <- struct{}{}

	// Wait for the end message from the goroutine.
	<-done
}
