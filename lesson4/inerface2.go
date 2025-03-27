package main

import "fmt"

// Empty interface can hold values of any type
func printValue(v interface{}) {
	fmt.Println("Value:", v)
}

func main() {
	printValue(42)           // Output: Value: 42
	printValue("Hello, Go!") // Output: Value: Hello, Go!
	printValue(3.14)         // Output: Value: 3.14
}