package zirreshte

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var strs []string
	for i := 0; i < n; i++ {
		var str string
		fmt.Scanln(&str)
		strs = append(strs, str)
	}

	lcs := longestCommonSubstring(strs)

	if lcs != "" {
		fmt.Println(lcs)
	}
}

func longestCommonSubstring(strs []string) string {

	n := len(strs)
	maxLen := 0
	for _, str := range strs {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}

	lcs := make([][]int, n)
	for i := 0; i < n; i++ {
		lcs[i] = make([]int, maxLen+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				lcs[i][j] = len(strs[i])
			} else {
				lcs[i][j] = 0
				for k := 0; k < len(strs[i]) && k < len(strs[j]); k++ {
					if strs[i][k] == strs[j][k] {
						lcs[i][j] = max(lcs[i][j], lcs[i-1][j-1]+1)
					}
				}
			}
		}
	}

	maxLength := 0
	lcsIdx := -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if lcs[i][j] > maxLength {
				maxLength = lcs[i][j]
				lcsIdx = i
			}
		}
	}

	if maxLength > 0 {
		return strs[lcsIdx][:maxLength]
	}

	return ""
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
