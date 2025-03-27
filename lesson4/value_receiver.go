package main

import "fmt"

// Define a struct for Author
type Author struct {
	name  string
	books []string
}

// Method with a value receiver to add a book
func (a Author) addBook(book string) {
	a.books = append(a.books, book)
}

func main() {
	author := Author{name: "J.K. Rowling", books: []string{"Harry Potter"}}
	author.addBook("Fantastic Beasts")
	fmt.Println("Books:", author.books) // Output: Books: [Harry Potter]
}