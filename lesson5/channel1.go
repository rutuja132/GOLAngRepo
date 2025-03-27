
package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // Send data to the channel
		}
		close(ch)
	}()

	// Receive data from the channel
	for value := range ch {
		fmt.Println("Received:", value)
	}
}
