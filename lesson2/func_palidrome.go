package main

import "fmt"

func isPalindrome(num int) bool {
	original, reversed := num, 0

	for num > 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}

	return original == reversed
}

func main() {
	var number int
	fmt.Print("Enter a number to check is palindrome or not : ")
	fmt.Scanln(&number)

	if isPalindrome(number) {
		fmt.Printf("%d is a palindrome.\n", number)
	} else {
		fmt.Printf("%d is not a palindrome number .\n", number)
	}
}