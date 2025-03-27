package main

import "fmt"

// Define a struct
type Rectangle struct {
	width, height int
}

// Method with a struct receiver
func (r Rectangle) area() int {
	return r.width * r.height
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	fmt.Println("Area of rectangle:", rect.area()) // Output: Area of rectangle: 50
}