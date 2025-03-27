package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Println("Sum:", num1+num2)