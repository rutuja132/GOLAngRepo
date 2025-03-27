package main

import (
	"fmt"
	"strconv"
	"sync"
)

func sumOfSquares(num int, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done()
	sum := 0
	for _, digit := range strconv.Itoa(num) {
		d := int(digit - '0') // Convert rune to int
		sum += d * d
	}
	result <- sum // Send the result to the channel
}

func sumOfCubes(num int, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done()
	sum := 0
	for _, digit := range strconv.Itoa(num) {
		d := int(digit - '0') // Convert rune to int
		sum += d * d * d
	}
	result <- sum // Send the result to the channel
}

func main() {
	num := 123
	var wg sync.WaitGroup
	result := make(chan int, 2) // Channel to receive results

	wg.Add(2) // We have two goroutines
	go sumOfSquares(num, &wg, result)
	go sumOfCubes(num, &wg, result)

	wg.Wait() // Wait for both goroutines to finish
	close(result)

	sumSquares := 0
	sumCubes := 0

	for res := range result {
		if sumSquares == 0 {
			sumSquares = res
		} else {
			sumCubes = res
		}
	}

	finalSum := sumSquares + sumCubes
	fmt.Printf("Sum of squares = %d\n", sumSquares)
	fmt.Printf("Sum of cubes = %d\n", sumCubes)
	fmt.Printf("Final sum of squares and cubes = %d\n", finalSum)
}