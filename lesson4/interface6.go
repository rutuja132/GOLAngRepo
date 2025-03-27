
package main

import "fmt"

// Define an interface
type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	var a Animal

	a = Dog{}
	fmt.Println(a.Speak()) // Output: Woof!

	a = Cat{}
	fmt.Println(a.Speak()) // Output: Meow!
}
