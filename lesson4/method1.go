package main

import "fmt"

// Type definition
type data int

// non-struct type receiver
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}
func main() {
	value1 := data(25)
	value2 := data(20)
	res := value1.multiply(value2)
	fmt.Println("Final result: ", res)
}