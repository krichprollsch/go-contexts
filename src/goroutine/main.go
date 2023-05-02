package main

import "fmt"

func main() {
	// Run a function synchronously.
	fmt.Print("Hello")

	// Run a function asynchronously, ie. in a goroutine.
	go fmt.Println(" world!")
}
