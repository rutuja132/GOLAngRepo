package main

import "fmt"

// base interface
type Vehicle interface {
	Start()
	Stop()
}

type Car interface {
	Vehicle // Embedding Vehicle interface
	Drive()
}

// Implement the interfaces in a struct
type Sedan struct {
	brand string
}

// Implement Vehicle methods for Sedan
func (s Sedan) Start() {
	fmt.Println(s.brand, "car is starting.")
}

func (s Sedan) Stop() {
	fmt.Println(s.brand, "car is stopping.")
}

// Implement the Car-specific method
func (s Sedan) Drive() {
	fmt.Println(s.brand, "car is driving.")
}

func main() {

	myCar := Sedan{brand: "Toyota"}

	myCar.Start()
	myCar.Drive()
	myCar.Stop()
}