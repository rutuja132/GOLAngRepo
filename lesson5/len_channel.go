
package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with capacity 3
	myChnlLen := make(chan string, 3)

	// Send values to the channel
	myChnlLen <- "A"
	myChnlLen <- "B"
	myChnlLen <- "C"

	// Print the length of the channel
	fmt.Println("Length of the channel (len):", len(myChnlLen)) // Should print 3
}
