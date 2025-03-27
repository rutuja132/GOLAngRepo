
package main

import "fmt"

type Student struct {
	RollNo   int
	StudName string
	Mark1    int
	Mark2    int
	Mark3    int
	Total    int
	Average  float64
}

func main() {
	var n int
	fmt.Print("Enter the number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {

		fmt.Printf("Enter details for student %d\n", i+1)
		fmt.Print("Roll No: ")
		fmt.Scan(&students[i].RollNo)
		fmt.Print("Name: ")
		fmt.Scan(&students[i].StudName)
		fmt.Print("Marks in 3 subjects (Mark1 Mark2 Mark3): ")
		fmt.Scan(&students[i].Mark1, &students[i].Mark2, &students[i].Mark3)

		students[i].Total = students[i].Mark1 + students[i].Mark2 + students[i].Mark3
		students[i].Average = float64(students[i].Total) / 3
	}

	fmt.Println("\nStudent Details:")
	for i := 0; i < n; i++ {
		fmt.Printf("Roll No: %d, Name: %s, Total: %d, Average: %.2f\n", students[i].RollNo, students[i].StudName, students[i].Total, students[i].Average)
	}
}
