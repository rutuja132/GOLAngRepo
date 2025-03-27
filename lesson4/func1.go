package main

import "fmt"

// Function to add two integers
func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(5, 3)
	fmt.Println("Sum:", result) // Output: Sum: 8
}