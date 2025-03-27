package main

import "fmt"

func identifyType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Type: int, Value:", v)
	case string:
		fmt.Println("Type: string, Value:", v)
	case bool:
		fmt.Println("Type: bool, Value:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	identifyType(42)
	identifyType("Hello, Go!")
	identifyType(true)
	identifyType(3.14)
}