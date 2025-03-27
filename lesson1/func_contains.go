package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Contains("test", "es")
	fmt.Println("Contains:", result)
}