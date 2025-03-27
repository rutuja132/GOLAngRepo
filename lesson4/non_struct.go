package main

import "fmt"

// Define a custom type
type MyInt int

// Method with a non-struct type receiver
func (m MyInt) double() MyInt {
	return m * 2
}

func main() {
	num := MyInt(5)
	fmt.Println("Double:", num.double()) // Output: Double: 10
}