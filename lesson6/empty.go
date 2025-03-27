package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("empty.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Empty file created successfully.")
}