package main

import "fmt"

// Define a struct to hold the numbers
type Calculator struct{}


func (c Calculator) Multiply(a int, b int) int {
    return a * b
}

func main() {
    calc := Calculator{}
    num1 := 5
    num2 := 10
    result := calc.Multiply(num1, num2)
    fmt.Printf("Multiplication of %d and %d is: %d\n", num1, num2, result)
}