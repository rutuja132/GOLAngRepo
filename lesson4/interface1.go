package main

import (
	"fmt"
)

type MyInterface interface{}

func main() {
	var value MyInterface

	value = 42
	fmt.Println("Integer Value:", value.(int))

	value = "Hello, Go!"
	fmt.Println("String Value:", value.(string))

	value = 3.14
	fmt.Println("Float Value:", value.(float64))

	value = true
	if booleanVal, ok := value.(bool); ok {
		fmt.Println("Boolean Value:", booleanVal)
	} else {
		fmt.Println("Type assertion failed")
	}
}