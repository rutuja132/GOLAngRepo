package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var waitGroup sync.WaitGroup // Create WaitGroup

	// Start 5 worker goroutines
	for i := 1; i <= 5; i++ {
		waitGroup.Add(1)         // Increment counter
		go worker(i, &waitGroup) // Launch goroutine
	}

	waitGroup.Wait() // Wait for all to finish
	fmt.Println("All workers completed.")
}