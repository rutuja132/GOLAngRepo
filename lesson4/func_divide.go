package main

import "fmt"

// Function that returns multiple values
func divide(a, b int) (int, int) {
	return a / b, a % b // Return quotient and remainder
}

func main() {
	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder) // Output: Quotient: 3 Remainder: 1
}