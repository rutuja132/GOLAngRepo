package main

import "fmt"

type ArrayCopier struct{}

func (ac ArrayCopier) CopyArray(source [5]int) [5]int {
	var destination [5]int
	for i, v := range source {
		destination[i] = v
	}
	return destination
}

func main() {
	source := [5]int{1, 2, 3, 4, 5}
	copier := ArrayCopier{}
	destination := copier.CopyArray(source)

	fmt.Println("Source Array:", source)
	fmt.Println("Destination Array:", destination)
}