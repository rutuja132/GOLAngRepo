package main

import "fmt"

func main() {
	
	fmt.Println("Using a basic for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}


	fmt.Println("\nUsing a loop with break:")
	count := 1
	for {
		if count > 3 {
			break
		}
		fmt.Println(count)
		count++
	}

	fmt.Println("\nUsing continue to skip even numbers:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue 
		}
		fmt.Println(i)
	}
}