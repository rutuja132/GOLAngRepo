package main

import "fmt"

// Function that modifies a value (call by reference)
func modifyValue(a *int) {
	*a = *a + 10
	fmt.Println("Inside modifyValue function:", *a)
}

func main() {
	num := 5
	fmt.Println("Before modifyValue:", num)
	modifyValue(&num)
	fmt.Println("After modifyValue:", num)
}