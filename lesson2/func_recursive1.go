package main

import "fmt"

// Indirect recursive function 1
func factorialHelper(n int) int {
	if n == 0 {
		return 1
	}
	return multiply(n, n-1)
}

func multiply(n, m int) int {
	return n * factorialHelper(m) // indirect recursion
}

func main() {
	var num int
	fmt.Print("Enter a number to calculate its factorial: ")
	fmt.Scan(&num)

	result := factorialHelper(num)
	fmt.Printf("Factorial of %d is %d\n", num, result)
}