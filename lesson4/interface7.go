package main

import "fmt"

type Shape interface {
	Area() float64
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

type Triangle struct {
	base, height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func main() {
	triangle := Triangle{base: 10, height: 5}
	printArea(triangle) // Output: Area: 25
}