package main

import (
	"fmt"
	"time"
)

func printNumbers(id int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

func main() {
	for i := 1; i <= 3; i++ {
		go printNumbers(i) // Start multiple goroutines
	}

	// Sleep to allow goroutines to finish
	time.Sleep(2 * time.Second)
	fmt.Println("Main function completed.")
}