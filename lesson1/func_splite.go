package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Split("a-b-c-d-e", "-")
	fmt.Println("Split:", result)
}