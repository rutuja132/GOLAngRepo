package main

import (
	"fmt"
	"sort"
)

type Student struct {
	rollNo     int
	name       string
	percentage float64
}

// Function to display student details
func displayStudents(students []Student) {
	fmt.Println("\nStudent Information in Descending Order of Percentage:")
	for _, s := range students {
		fmt.Printf("Roll No: %d, Name: %s, Percentage: %.2f%%\n", s.rollNo, s.name, s.percentage)
	}
}

func main() {
	var n int

	fmt.Print("Enter the number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Student %d:\n", i+1)
		fmt.Print("Roll No: ")
		fmt.Scan(&students[i].rollNo)
		fmt.Print("Name: ")
		fmt.Scan(&students[i].name)
		fmt.Print("Percentage: ")
		fmt.Scan(&students[i].percentage)
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].percentage > students[j].percentage
	})

	displayStudents(students)
}