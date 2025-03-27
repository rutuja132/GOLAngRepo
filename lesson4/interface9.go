package main

import "fmt"

type Stringer interface {
	String() string
}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.name, p.age)
}

func main() {
	var s Stringer = Person{name: "Alice", age: 30}
	fmt.Println(s.String()) // Output: Alice is 30 years old
}

.