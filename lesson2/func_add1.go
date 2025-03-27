package main

import "fmt"

// Function to add two numbers
func add(a int, b int) int {
	return a + b
}

func main() {
	var num1, num2 int

	// Taking input from the user
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Calling the add function and printing the result
	result := add(num1, num2)
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, result)
}