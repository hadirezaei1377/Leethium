package paranteza

import (
	"fmt"
)

func main() {
	var s string

	fmt.Scan(&s)

	stack := make([]byte, 0)
	imbalance := 0

	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' {
			if len(stack) == 0 {
				imbalance += 1
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	imbalance += len(stack) / 2

	fmt.Println(imbalance)
}
