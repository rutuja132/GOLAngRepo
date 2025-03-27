package main

import "fmt"

func main() {
	var theArray [3]string
	theArray[0] = "India"
	theArray[1] = "Canada" // Assign a value to the second element
	theArray[2] = "USA"
	fmt.Println(theArray[0])
	fmt.Println(theArray[1]) // Access the second element valu
	fmt.Println(theArray[2])
}