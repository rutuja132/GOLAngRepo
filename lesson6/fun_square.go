package main

import "fmt"

// Square function to find the square of a number
func Square(n int) int {
	return n * n
}

func main() {
	num := 4
	result := Square(num)
	fmt.Printf("The square of %d is %d\n", num, result)
}