package main

import (
	"fmt"
	"sync"
)

func checkEven(num int, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done()
	if num%2 == 0 {
		result <- num // Send even number to channel
	}
}

func checkOdd(num int, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done()
	if num%2 != 0 {
		result <- num // Send odd number to channel
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup
	result := make(chan int)

	// Start goroutines to check even and odd numbers
	for _, num := range numbers {
		wg.Add(1)
		if num%2 == 0 {
			go checkEven(num, &wg, result)
		} else {
			go checkOdd(num, &wg, result)
		}
	}

	go func() {
		wg.Wait()     // Wait for all goroutines to finish
		close(result) // Close the channel when done
	}()

	// Receive results from the channel
	for res := range result {
		fmt.Println("Received:", res) // Print received values
	}
}