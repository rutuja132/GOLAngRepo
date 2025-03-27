package main

import (
	"fmt"
)

func main() {

	myChnl := make(chan string)

	go func() {
		for _, msg := range []string{"Hello", "World", "Go", "Programming"} {
			myChnl <- msg // Send message to channel
		}
		close(myChnl)
	}()

	// Receive values from the channel using for range
	for msg := range myChnl {
		fmt.Println("Received:", msg)
	}
}