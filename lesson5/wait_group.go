package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go f1(wg)
	go f2(wg)
	fmt.Print("\n Hi ! from main go routine")
	wg.Wait()
}
func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\n In f1 go routine")
}
func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\n In f2 go routine")
}