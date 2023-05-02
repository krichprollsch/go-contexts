package main

import (
	"fmt"
	"time"
)

// slow simulates a slow operation, like an HTTP request.
func slow() {
	time.Sleep(2 * time.Second)
	fmt.Println(" world!")
}

func main() {
	// Run a function synchronously.
	fmt.Print("Hello")

	// done is a channel used to synchronize the goroutine and the main program.
	done := make(chan struct{})

	// Run a function asynchronously, ie. in a goroutine.
	go func() {
		slow()
		// Once the function stops, we send a message in the channel.
		done <- struct{}{}
	}()

	// Wait for the end message from the goroutine.
	<-done
}
