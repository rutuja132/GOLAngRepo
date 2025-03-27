package main

import (
	"calculator"
	"fmt"
)

func main() {
	var a, b float64
	var choice int

	fmt.Println("Enter two numbers:")
	fmt.Scan(&a, &b)

	fmt.Println("Choose an operation:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("Result: %.2f\n", calculator.Add(a, b))
	case 2:
		fmt.Printf("Result: %.2f\n", calculator.Subtract(a, b))
	case 3:
		fmt.Printf("Result: %.2f\n", calculator.Multiply(a, b))
	case 4:
		fmt.Printf("Result: %.2f\n", calculator.Divide(a, b))
	default:
		fmt.Println("Invalid choice")
	}
}