package main

import "fmt"

func square(num int) int {
	return num * num
}

func main() {
	result := square(4)
	fmt.Printf("Square of 4 is: %d\n", result)
}