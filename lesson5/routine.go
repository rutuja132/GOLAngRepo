package main

import "fmt"

func main() {
	message()
	fmt.Print("Hi ! from main go routine")
}
func message() {
	fmt.Print("Hi ! from message go routine")
}