package main

import (
	"context"
	"fmt"
	"time"
)

// slow simulates a slow operation, like an HTTP request.
func slow(ctx context.Context) {
	select {
	// Simulate a long operation, like an HTTP request.
	case <-time.After(2 * time.Second):
		// nothing to do here
	case <-ctx.Done():
		// operation is canceled, return early.
		return
	}

	fmt.Println(" world!")
}

func main() {
	// create a cancelable context for our main program.
	ctx, cancel := context.WithCancel(context.Background())

	// Run a function synchronously.
	fmt.Print("Hello")

	// done is a channel used to synchronize the goroutine and the main program.
	done := make(chan struct{})

	// Run a function asynchronously, ie. in a goroutine.
	go func() {
		slow(ctx)
		// Once the function stops, we send a message in the channel.
		done <- struct{}{}
	}()

	// Cancel the context.
	cancel()

	// Wait for the end message from the goroutine.
	<-done
}
