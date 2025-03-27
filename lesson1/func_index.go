package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Index("test", "e")
	fmt.Println("Index:", result)
}