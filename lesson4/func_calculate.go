package main

import "fmt"

// Function with named return values
func calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // Return named values
}

func main() {
	s, p := calculate(4, 5)
	fmt.Println("Sum:", s, "Product:", p) // Output: Sum: 9 Product: 20
}