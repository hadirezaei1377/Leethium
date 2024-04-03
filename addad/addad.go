package addad

import (
	"fmt"
	"strconv"
)

func main() {
	var k int
	fmt.Println("Enter k:")
	fmt.Scan(&k)

	numStr := ""
	for i := 1; i <= 5000; i++ {
		numStr += strconv.Itoa(i)
	}

	if k > 0 && k <= len(numStr) {
		fmt.Printf("The %dth digit from the left is: %c\n", k, numStr[k-1])
	} else {
		fmt.Println("Invalid input for k!")
	}
}
