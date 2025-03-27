package main

import "fmt"

func main() {
	var input string

	// Prompting the user for input
	fmt.Println("Enter a color (red, blue, green):")

	// Reading user input
	fmt.Scanln(&input)

	// Using switch to handle different input values
	switch input {
	case "red":
		fmt.Println("You chose Red!")
	case "blue":
		fmt.Println("You chose Blue!")
	case "green":
		fmt.Println("You chose Green!")
	default:
		fmt.Println("Invalid color chosen.")
	}
}