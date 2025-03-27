package main

import "fmt"

func swapNumbers(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	var num1, num2 int
	fmt.Print("Enter number to swap numbers ")

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Printf("Before swapping: num1 = %d, num2 = %d\n", num1, num2)

	swapNumbers(&num1, &num2)

	fmt.Printf("After swapping: num1 = %d, num2 = %d\n", num1, num2)
	
}