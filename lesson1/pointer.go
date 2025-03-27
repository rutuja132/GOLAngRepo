
package main

import "fmt"

func main() {

	var num int = 10

	fmt.Println("Value of num:", num)

	var ptr *int = &num

	fmt.Println("Value via pointer ptr:", *ptr)

	var ptr2 **int = &ptr

	fmt.Println("Value via pointer to pointer ptr2:", **ptr2)
}
