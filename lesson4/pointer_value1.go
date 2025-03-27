package main

import "fmt"

// Define a struct
type Point struct {
	x, y int
}

// Method with a value receiver
func (p Point) move(dx, dy int) {
	p.x += dx
	p.y += dy
}

// Method with a pointer receiver
func (p *Point) movePointer(dx, dy int) {
	p.x += dx
	p.y += dy
}

func main() {
	p1 := Point{x: 1, y: 2}
	p1.move(1, 1)                              // Value receiver, p1 remains unchanged
	fmt.Println("Point after value move:", p1) // Output: Point after value move: {1 2}

	p1.movePointer(1, 1)                         // Pointer receiver, p1 is modified
	fmt.Println("Point after pointer move:", p1) // Output: Point after pointer move: {2 3}
}