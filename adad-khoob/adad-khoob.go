package adadkhoob

import (
	"fmt"
	"math"
)

func countDivisors(n int) int {
	count := 0
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			count += 2
			if i == n/i {
				count--
			}
		}
	}
	return count
}

func main() {
	var k int
	fmt.Scan(&k)

	n := 1
	for {

		goodNumber := n * (n + 1) / 2

		if countDivisors(goodNumber) >= k {
			fmt.Println(goodNumber)
			break
		}

		n++
	}
}
