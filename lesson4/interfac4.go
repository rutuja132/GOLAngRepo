
package main

import "fmt"

type Counter interface {
	Increment()
	Value() int
}

type SimpleCounter struct {
	count int
}

func (c *SimpleCounter) Increment() {
	c.count++
}

func (c *SimpleCounter) Value() int {
	return c.count
}

func main() {
	var c Counter = &SimpleCounter{}
	c.Increment()
	c.Increment()
	fmt.Println("Counter value:", c.Value()) // Output: Counter value: 2
}
