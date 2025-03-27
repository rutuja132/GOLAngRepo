package main

import (
	"fmt"
	"sync"
)

func sum(numbers []int, wg *sync.WaitGroup, result *int) {
	defer wg.Done() // Signal completion
	for _, number := range numbers {
		*result += number
	}
}

func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5}
	var result int

	// Split the work into two goroutines
	wg.Add(2)                         // Two goroutines
	go sum(numbers[:3], &wg, &result) // First half
	go sum(numbers[3:], &wg, &result) // Second half

	wg.Wait() // Wait for the - all to finish
	fmt.Printf("Total sum: %d\n", result)
}