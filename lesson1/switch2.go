package main

import "fmt"

func main() {
	var day int = 3

	switch day {
	case 1, 2, 3:
		fmt.Println("Start of the week")
	case 4, 5:
		fmt.Println("Midweek")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}
}