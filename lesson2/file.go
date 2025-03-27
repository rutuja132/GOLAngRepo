package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a file
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Defer closing the file
	defer file.Close()

	// Write "Hello World" to the file
	_, err = file.WriteString("Hello World")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("File created and written successfully")
	}
}