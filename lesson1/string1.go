package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter your name and age:")
	fmt.Scan(&name, &age)

	fmt.Printf("Hello %s, you are %d years old.\n", name, age)
}