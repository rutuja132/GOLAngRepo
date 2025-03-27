
package main

import "fmt"

func calculateRectangleProperties(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

func main() {
	var length, width float64
	fmt.Print("Enter length of the rectangle: ")
	fmt.Scanln(&length)
	fmt.Print("Enter width of the rectangle: ")
	fmt.Scanln(&width)

	area, perimeter := calculateRectangleProperties(length, width)
	fmt.Printf("Area: %.2f\nPerimeter: %.2f\n", area, perimeter)
}
