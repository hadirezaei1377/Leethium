package ascendingsorting

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)

	permutation := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&permutation[i])
	}

	canReverseSort := canReverseSortWithZeros(n, permutation)

	if canReverseSort {

		reversePermutation, replacements := findReverseSortPermutation(n, permutation)
		fmt.Println("Yes")
		fmt.Println(reversePermutation)
		fmt.Println(replacements)
	} else {
		fmt.Println("No")
	}
}

func canReverseSortWithZeros(n int, permutation []int) bool {
	zeroCount := 0
	for _, num := range permutation {
		if num == 0 {
			zeroCount++
		}
	}

	usedNumbers := make(map[int]bool)
	for _, num := range permutation {
		if num != 0 && usedNumbers[num] {
			return false
		}
		usedNumbers[num] = true
	}

	for _, num := range permutation {
		if num > n {
			return false
		}
	}

	if zeroCount < n-1 {
		return false
	}

	return true
}

func findReverseSortPermutation(n int, permutation []int) ([]int, string) {
	usedNumbers := make(map[int]bool)
	for i, num := range permutation {
		if num == 0 {
			for j := 1; j <= n; j++ {
				if !usedNumbers[j] {
					permutation[i] = j
					usedNumbers[j] = true
					break
				}
			}
		}
	}

	sort.Slice(permutation, func(i, j int) bool {
		return permutation[j] > permutation[i]
	})

	replacements := ""
	for i, num := range permutation {
		if i > 0 && permutation[i] < permutation[i-1] {
			replacements += fmt.Sprintf("%d=%d ", permutation[i-1], num)
		}
	}

	return permutation, replacements
}
