package main

import "fmt"

func main() {
	var num int

	fmt.Println("type  a number to check is single digit or not :")
	fmt.Scan(&num)

	if num >= -9 && num <= 9 {
		fmt.Println("The number is a single digit.")
	} else {
		fmt.Println(" is not a single digit.")
	}
}