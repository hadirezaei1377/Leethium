package jenasghalb

import "fmt"

func isPalindrome(s string) bool {

	n := len(s)

	if n <= 1 {
		return true
	}

	if s[0] != s[n-1] {
		return false
	}

	return isPalindrome(s[1 : n-1])
}

func main() {

	inputs := []string{"Radar", "a man. a plan? a canal Panama", "Hi h", "Book", "0120"}

	for _, input := range inputs {
		result := "NO"
		if isPalindrome(input) {
			result = "YES"
		}
		fmt.Println("input:", input)
		fmt.Println("output:", result)
	}
}
