package main

import "fmt"

// Subtract function to subtract two integers
func Subtract(a int, b int) int {
	return a - b
}

func main() {
	a, b := 10, 5
	result := Subtract(a, b)
	fmt.Printf("The result of %d - %d is %d\n", a, b, result)
}