package main

import "fmt"

func main() {
	// Run a function synchronously.
	fmt.Print("Hello")

	// done is a channel used to synchronize the goroutine and the main program.
	done := make(chan struct{})

	// Run a function asynchronously, ie. in a goroutine.
	go func() {
		fmt.Println(" world!")
		// Once the function stops, we send a message in the channel.
		done <- struct{}{}
	}()

	// Wait for the end message from the goroutine.
	<-done
}
