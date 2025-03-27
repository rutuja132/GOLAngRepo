package main

import "fmt"

func multi(a, b int) int {
	var r int
	r = a * b
	return r
}
func main() {
	c := multi(2, 4)
	fmt.Println("multiplication is ", c)
}