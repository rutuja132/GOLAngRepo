package main

import (
	"fmt"
)

func main() {
	// Create an unbuffered channel
	myChnlClose := make(chan string)

	// Start a goroutine to send values
	go func() {
		myChnlClose <- "Hello"
		myChnlClose <- "World"
		close(myChnlClose) // Close the channel
	}()

	// Receive values until the channel is closed
	for msg := range myChnlClose {
		fmt.Println("Received (close):", msg) // Print received messages
	}
}