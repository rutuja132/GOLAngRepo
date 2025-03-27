package main

import "fmt"

// Function that returns multiple values
func calculate(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func main() {
	var num1, num2 int
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	sum, product := calculate(num1, num2)
	fmt.Printf("Sum: %d\nProduct: %d\n", sum, product)
}