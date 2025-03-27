package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3) // Create a buffered channel with a capacity of 3

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Sending value:", i)
			ch <- i // Send value to the channel
			fmt.Println("Value sent:", i)
		}
		close(ch)
	}()

	// Receive values from the channel
	for value := range ch {
		fmt.Println("Received:", value) // Print received values
	}
}