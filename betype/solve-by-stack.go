package betype

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	stack := []rune{}

	for _, char := range input {
		if char == '=' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, char)
		}
	}

	result := string(stack)
	fmt.Println(result)
}
