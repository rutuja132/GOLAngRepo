package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func getShape() Shape {
	return Square{side: 4}
}

func main() {
	shape := getShape()
	fmt.Println("Area of shape:", shape.Area()) // Output: Area of shape: 16
}