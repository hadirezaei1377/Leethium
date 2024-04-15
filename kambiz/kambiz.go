package kamviz

import (
	"fmt"
	"strconv"
)

func printValidIPs(s string) {
	for i := 1; i <= 3 && i < len(s); i++ {
		for j := i + 1; j <= i+3 && j < len(s); j++ {
			for k := j + 1; k <= j+3 && k < len(s); k++ {
				part1, _ := strconv.Atoi(s[:i])
				part2, _ := strconv.Atoi(s[i:j])
				part3, _ := strconv.Atoi(s[j:k])
				part4, _ := strconv.Atoi(s[k:])

				if isValid(part1) && isValid(part2) && isValid(part3) && isValid(part4) {

					fmt.Printf("%d.%d.%d.%d\n", part1, part2, part3, part4)
				}
			}
		}
	}
}

func isValid(part int) bool {
	return part >= 0 && part <= 255
}

func main() {
	var s string
	fmt.Scanln(&s)
	printValidIPs(s)
}
