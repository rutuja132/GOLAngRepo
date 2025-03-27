package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of elements in the array: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	largest, smallest := arr[0], arr[0]

	for _, num := range arr {
		if num > largest {
			largest = num
		}
		if num < smallest {
			smallest = num
		}
	}

	fmt.Printf("Largest number: %d\n", largest)
	fmt.Printf("Smallest number: %d\n", smallest)
}