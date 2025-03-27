package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("The number", number, "is Even.")
	} else {
		fmt.Println("The number", number, "is Odd.")
	}
}