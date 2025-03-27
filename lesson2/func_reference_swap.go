package main

import "fmt"

// Function that swaps two numbers (call by reference)
func swap(a, b *int) {
	*a, *b = *b, *a
	fmt.Printf("Inside swap function: a = %d, b = %d\n", *a, *b)
}

func main() {
	x, y := 5, 10
	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap: x = %d, y = %d\n", x, y)
}