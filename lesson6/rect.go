
package main

import (
	"fmt"
	"rectangle"
)

func main() {
	var length, width float64

	fmt.Println("Enter length and width of the rectangle:")
	fmt.Scan(&length, &width)

	area := rectangle.Area(length, width)
	fmt.Printf("Area of the rectangle: %.2f\n", area)
}
