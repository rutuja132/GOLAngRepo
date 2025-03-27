package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	fmt.Print("Enter the filename: ")
	fmt.Scan(&filename)

	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size(), "bytes")
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("Permissions:", fileInfo.Mode())
}