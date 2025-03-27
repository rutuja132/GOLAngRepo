package main

import (
	"fmt"
)

// Function type that takes two integers and returns an integer
type operation func(int, int) int

// Function that takes two integers and a function as parameters
func calculate(a int, b int, op operation) int {
	return op(a, b) // Call the passed function with a and b
}

// Sample functions to be passed as parameters
func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func main() {
	a, b := 10, 5

	// Using the calculate function with different operations
	sum := calculate(a, b, add)
	fmt.Printf("Sum: %d\n", sum) // Output: Sum: 15

	difference := calculate(a, b, subtract)
	fmt.Printf("Difference: %d\n", difference) // Output: Difference: 5

	product := calculate(a, b, multiply)
	fmt.Printf("Product: %d\n", product) // Output: Product: 50
}