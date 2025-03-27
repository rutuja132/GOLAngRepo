package main

import (
	"fmt"
)

func main() {

	myChnlCap := make(chan string, 5)

	// Send values to the channel
	myChnlCap <- "X"
	myChnlCap <- "Y"
	myChnlCap <- "Z"

	fmt.Println("Capacity of the channel (cap):", cap(myChnlCap)) // Should print 5
}