package main

import (
	"fmt"
	"scan"
)

func main() {
	var n int
	scan.Scan(&n)
	heights := make([]int, n)
	for i := 0; i < n; i++ {
		scan.Scan(&heights[i])
	}

	maxFlowers := 0

	for i := 0; i < n; i++ {
		maxHeight := heights[i]
		currentFlowers := 1

		for j := i + 1; j < n; j++ {
			if heights[j] >= maxHeight {
				maxHeight = heights[j]
				currentFlowers++
			}
		}

		if currentFlowers > maxFlowers {
			maxFlowers = currentFlowers
		}
	}

	fmt.Println(maxFlowers)
}
