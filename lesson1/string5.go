package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Println("Enter the first string to check substring or not:")
	fmt.Scan(&str1)
	fmt.Println("Enter the second string to check:")
	fmt.Scan(&str2)

	if strings.Contains(str2, str1) {
		fmt.Println("The first string is a substring of the second string.")
	} else {
		fmt.Println("The first string is not a substring of the second string.")
	}
}