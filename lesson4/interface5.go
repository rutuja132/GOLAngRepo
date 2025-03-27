package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	shapes := []Shape{
		Circle{radius: 5},
		Rectangle{width: 10, height: 5},
	}

	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
	}
}