package main

import "fmt"

// Define a struct
type Number struct {
	value int
}

// Method that can accept both pointer and value receivers
func (n Number) getValue() int {
	return n.value
}

func (n *Number) setValue(val int) {
	n.value = val
}

func main() {
	num := Number{value: 10}
	fmt.Println("Value:", num.getValue()) // Output: Value: 10

	num.setValue(20)
	fmt.Println("Updated Value:", num.getValue()) // Output: Updated Value: 20
}