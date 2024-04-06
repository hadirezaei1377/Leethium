package main

import (
	"fmt"
	"strings"
)

func convertToBase(num, base int) string {
	if num == 0 {
		return "0"
	}

	digits := "0123456789ABCDEF"
	var result strings.Builder

	for num > 0 {
		remainder := num % base
		result.WriteString(string(digits[remainder]))
		num /= base
	}

	// Reverse the string
	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	var num, base int
	fmt.Print("Enter the number in base 10: ")
	fmt.Scan(&num)
	fmt.Print("Enter the desired base: ")
	fmt.Scan(&base)

	result := convertToBase(num, base)
	fmt.Println("Result:", result)
}
