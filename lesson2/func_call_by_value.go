package main

import "fmt"

func modifyValue(num int) {
	num = 100
}

func main() {
	var number int = 50
	fmt.Printf("Before function call: number = %d\n", number)

	modifyValue(number)

	fmt.Printf("After function call: number = %d (unchanged)\n", number)
}