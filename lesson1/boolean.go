package main

import "fmt"

func main() {
	var isRaining bool = true

	switch isRaining {
	case true:
		fmt.Println("It is raining.")
	case false:
		fmt.Println("It is not raining.")
	default:
		fmt.Println("Unknown weather condition.")
	}
}