
package main

import "fmt"

func main() {
	var str1, str2 string

	fmt.Println("Enter 2 strings to compare :")
	fmt.Scan(&str1, &str2)

	if str1 == str2 {
		fmt.Println("The strings are equal.")
	} else {
		fmt.Println("The strings are not equal.")
	}
}
