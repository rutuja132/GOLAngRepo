
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the number of terms in the Fibonacci series: ")
	fmt.Scan(&n)

	fmt.Println("Fibonacci Series:")
	fibonacci(n)
}

func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}
