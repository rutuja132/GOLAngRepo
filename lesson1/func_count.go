package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Count("test", "t")
	fmt.Println("Count:", result)
}