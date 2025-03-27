package main
import (
	"fmt"
	"math"
)
type Shape interface{
	Area() float64
	Perimeter() float64
}

type Circle struct{
	
}