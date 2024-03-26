package khodmaghloob

import "fmt"

func reverseNumber(num int) int {
	reversed := 0
	for num > 0 {
		remainder := num % 10
		reversed = reversed*10 + remainder
		num /= 10
	}
	return reversed
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	original := num
	reversed := reverseNumber(num)

	if original == reversed {
		fmt.Println("The number is a khodmaghloob number")
	} else {
		fmt.Println("The number is not a khodmaghloob number")
	}
}

// 2356532 is a khodmaghloob number
// 4588 is not a khodmaghloob number
