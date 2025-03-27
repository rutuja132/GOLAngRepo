
package main

import "fmt"

// Define a structure
type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	// Initialize a struct (Method 1: using field names)
	person1 := Person{Name: "Alice", Age: 25, City: "New York"}

	// Initialize a struct (Method 2: without field names, order matters)
	person2 := Person{"Bob", 30, "Los Angeles"}

	// Initialize an empty struct (fields get default values)
	var person3 Person

	// Display the structs
	fmt.Println("Person 1:", person1)
	fmt.Println("Person 2:", person2)
	fmt.Println("Person 3 (default values):", person3)
}
