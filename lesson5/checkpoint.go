package main

import (
	"fmt"
	"sync"
	"time"
)

const numWorkers = 5

func worker(id int, wg *sync.WaitGroup, checkpoint chan struct{}) {
	defer wg.Done()

	workDuration := time.Duration(1+id) * time.Second
	fmt.Printf("Worker %d is working for %v seconds...\n", id, workDuration.Seconds())
	time.Sleep(workDuration)

	fmt.Printf("Worker %d has completed its task and is waiting at the checkpoint.\n", id)
	checkpoint <- struct{}{} // Signal that this worker has reached the checkpoint

	<-checkpoint
	fmt.Printf("Worker %d is proceeding to the next task.\n", id)
}

func main() {
	var wg sync.WaitGroup
	checkpoint := make(chan struct{}, numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, checkpoint)
	}

	wg.Wait()
	fmt.Println("All workers have completed their tasks.")
}