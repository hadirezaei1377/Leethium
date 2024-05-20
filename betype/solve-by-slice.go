package betype

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var result []rune

	for _, char := range input {
		if char == '=' {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, char)
		}
	}

	fmt.Println(string(result))
}
