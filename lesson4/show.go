
package main

import (
	"fmt"
)

type Author struct {
	Name  string
	Age   int
	Email string
}

func (a Author) Show() {
	fmt.Printf("Author Name: %s\n", a.Name)
	fmt.Printf("Author Age: %d\n", a.Age)
	fmt.Printf("Author Email: %s\n", a.Email)
}

func main() {

	author := Author{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@example.com",
	}

	author.Show()
}
