package main

import "fmt"

func main() {
	// Looping over a map
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}

	for name, age := range ages {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
}