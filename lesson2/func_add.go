package main

import "fmt"

func abc(n1, n2 int) int {
	var r int
	r = n1 + n2
	return r
}

func main() {
	result := abc(5, 7)
	fmt.Println("The sum is:", result)
}