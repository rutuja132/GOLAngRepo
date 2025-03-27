package main

import (
	"fmt"
)

func main() {
	// Example array
	arr := []int{12, 45, 2, 10, 33, 89, 4}

	// Initialize variables
	largest := arr[0]
	smallest := arr[0]

	// Iterate through the array
	for _, num := range arr {
		if num > largest {
			largest = num
		}
		if num < smallest {
			smallest = num
		}
	}

	// Print the results
	fmt.Println("Largest number:", largest)
	fmt.Println("Smallest number:", smallest)
}