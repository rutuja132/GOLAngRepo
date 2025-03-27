package main

import "fmt"

func main() {
	var m int

	fmt.Println("enter number to create table ")
	fmt.Print("\t")
	fmt.Scan(&m)
	fmt.Println("table of ", m)

	for i := 1; i <= 10; i++ {
		fmt.Println(m * i)
	}
}