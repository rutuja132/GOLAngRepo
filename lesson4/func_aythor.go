package main

import "fmt"

// Define a struct for Author
type Author struct {
	name  string
	books []string
}

// Method to display author details
func (a Author) details() {
	fmt.Printf("Author: %s\nBooks: %v\n", a.name, a.books)
}

func main() {
	author := Author{name: "George Orwell", books: []string{"1984", "Animal Farm"}}
	author.details() // Output: Author: George Orwell
}