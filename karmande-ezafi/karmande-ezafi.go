package karmandeezafi

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	nameCount := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		var firstName, lastName string
		fmt.Sscanf(scanner.Text(), "%s %s", &firstName, &lastName)
		nameCount[firstName]++
	}

	maxColors := 0
	for _, count := range nameCount {
		if count > maxColors {
			maxColors = count
		}
	}

	fmt.Println(maxColors)
}
