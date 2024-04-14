package jenas

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var input string
	fmt.Scanln(&input)

	if IsPalindrome(input) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}, s)

	return isPalindromeRecursive(s)
}

func isPalindromeRecursive(s string) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return isPalindromeRecursive(s[1 : len(s)-1])
}
