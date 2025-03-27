package main

import "fmt"

type Printer interface {
	Print()
}

type Message struct {
	content string
}

func (m Message) Print() {
	fmt.Println(m.content)
}

func printMessage(p Printer) {
	p.Print()
}

func main() {
	msg := Message{content: "Hello, Interfaces!"}
	printMessage(msg) // Output: Hello, Interfaces!
}