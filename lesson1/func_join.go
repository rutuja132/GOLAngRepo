package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Join([]string{"a", "b"}, "-")
	fmt.Println("Join:", result)
}