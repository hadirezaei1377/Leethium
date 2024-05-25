package zirmajmooe

import (
	"fmt"
	"strconv"
	"strings"
)

func generateSubsets(n int, index int, current []int) {
	if index > n {

		var builder strings.Builder
		builder.WriteString("{ ")
		for _, v := range current {
			builder.WriteString(strconv.Itoa(v))
			builder.WriteString(" ")
		}
		builder.WriteString("}")
		fmt.Println(builder.String())
		return
	}

	generateSubsets(n, index+1, append(current, index))

	generateSubsets(n, index+1, current)
}

func main() {
	var n int
	fmt.Scan(&n)
	generateSubsets(n, 1, []int{})
}
