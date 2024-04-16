package main

import (
	"fmt"
)

func maxPouzooFans(s string) int {
	maxFans := 0
	currentFans := 0

	for _, char := range s {
		if char == '0' {
			currentFans++
		} else {
			if currentFans > maxFans {
				maxFans = currentFans
			}
			currentFans = 0
		}
	}

	if currentFans > maxFans {
		maxFans = currentFans
	}

	return maxFans
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(maxPouzooFans(input))
}
