package ghaem

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Enter three positive integers:")

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	sides := []int{a, b, c}

	sort.Ints(sides)

	if sides[0]*sides[0]+sides[1]*sides[1] == sides[2]*sides[2] {

		fmt.Println("YES")
	} else {

		fmt.Println("NO")
	}
}
