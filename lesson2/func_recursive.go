package main

import "fmt"

// Direct recursive function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1 // base case
	}
	return n * factorial(n-1)
}

func main() {
	var num int
	fmt.Print("Enter a number to calculate its factorial: ")
	fmt.Scan(&num)

	result := factorial(num)
	fmt.Printf("Factorial of %d is %d\n", num, result)
}