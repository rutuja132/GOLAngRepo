package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Replace("foo", "o", "0", -1)
	fmt.Println("Replace:", result)
}