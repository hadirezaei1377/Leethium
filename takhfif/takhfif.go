package takhfif

import (
	"fmt"
	"unicode"
)

func isValidDiscountCode(code string) bool {

	subtitleMap := make(map[rune]bool)

	for _, char := range code {

		if unicode.IsLetter(char) {

			subtitleMap[unicode.ToLower(char)] = true
		}
	}

	return len(subtitleMap) == 6
}

func main() {

	discountCodes := []string{"quera102", "quEra0012", "qu0erraa12", "sN0Ap12", "qurra00L"}

	for _, code := range discountCodes {

		if isValidDiscountCode(code) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
