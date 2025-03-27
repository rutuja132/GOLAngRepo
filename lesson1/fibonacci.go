package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter the number forFibonacci series:")
	fmt.Scan(&n)

	a, b := 0, 1

	fmt.Println("Fibonacci series:")
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}