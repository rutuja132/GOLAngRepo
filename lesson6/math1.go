
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateRandomNumber generates a random number between 0 and max
func GenerateRandomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// PrintRandomNumber prints a random number
func PrintRandomNumber() {
	randomNumber := GenerateRandomNumber(100)
	fmt.Printf("Random number: %d\n", randomNumber)
}

func main() {
	PrintRandomNumber()
	PrintRandomNumber()
	PrintRandomNumber()
}
