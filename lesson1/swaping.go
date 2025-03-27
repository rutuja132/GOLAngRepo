package main

import "fmt"

func main() {
	var a, b, temp int

	fmt.Println("enter a two numbers to swap ")
	fmt.Scan(&a, &b)

	fmt.Println("numbers before swapping ")
	fmt.Println("a = ", a, "b = ", b)

	temp = a
	a = b
	b = temp

	fmt.Println("numbers after swapping ")
	fmt.Println("a = ", a, "b = ", b)
}