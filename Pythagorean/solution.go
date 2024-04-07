package main

import (
	"fmt"
)

func findPythagoreanTriple(n int) []int {
	for a := 1; a < n; a++ {
		for b := a; b < n; b++ {
			c := n - a - b
			if c < b {
				break
			}
			if a*a+b*b == c*c {
				return []int{a, b, c}
			}
		}
	}
	return nil
}

func main() {
	var n int
	fmt.Scan(&n)

	result := findPythagoreanTriple(n)

	if result != nil {
		for i := 0; i < len(result); i++ {
			fmt.Print(result[i], " ")
		}
	} else {
		fmt.Println("Impossible")
	}
}
