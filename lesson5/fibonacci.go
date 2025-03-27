package main

import (
	"fmt"
)

func fibonacci(n int, ch chan<- int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a // Send Fibonacci number to channel
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	n := 10
	ch := make(chan int)

	go fibonacci(n, ch)

	for num := range ch {
		fmt.Println(num)
	}
}