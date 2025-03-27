package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateNumbers(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		waitTime := rand.Intn(250)
		time.Sleep(time.Duration(waitTime) * time.Millisecond)
		fmt.Printf("Goroutine %d: %d\n", id, i)
	}
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go generateNumbers(i, &wg)
	}

	wg.Wait()
}