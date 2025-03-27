package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	var a Animal = Dog{}

	// Type assertion
	if dog, ok := a.(Dog); ok {
		fmt.Println("Dog says:", dog.Speak()) // Output: Dog says: Woof!
	} else {
		fmt.Println("Not a Dog")
	}
}