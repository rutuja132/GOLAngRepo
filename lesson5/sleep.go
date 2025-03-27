package main

import (
	"fmt"
	"time"
)

func main() {
	go f1()
	go f2()
	fmt.Print("\n Hi ! from main go routine")
	time.Sleep(time.Second)
}
func f1() {
	fmt.Print("\n In f1 go routine")
}
func f2() {
	fmt.Print("\n In f2 go routine")
}