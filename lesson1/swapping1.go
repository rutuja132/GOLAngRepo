package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter the value of a: ")
	fmt.Scan(&a)
	fmt.Print("Enter the value of b: ")
	fmt.Scan(&b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After swapping:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}