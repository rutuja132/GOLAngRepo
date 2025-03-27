package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start a goroutine to send messages to ch1
	go func() {
		time.Sleep(1 * time.Second) // Simulate some work
		ch1 <- "Message from channel 1"
	}()

	// Start a goroutine to send messages to ch2
	go func() {
		time.Sleep(2 * time.Second) // Simulate some work
		ch2 <- "Message from channel 2"
	}()

	// Use select to wait for messages from either channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1) // Handle message from ch1
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2) // Handle message from ch2
		}
	}

	fmt.Println("All messages received.")
}