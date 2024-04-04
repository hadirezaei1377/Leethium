package majid

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	colors := make(map[int]int)
	for i := 0; i < n; i++ {
		var color int
		fmt.Scan(&color)
		colors[color]++
	}

	minColor := findMinColor(colors)
	fmt.Println(minColor)
}

func findMinColor(colors map[int]int) int {
	var minColor int
	minCount := len(colors) + 1

	for color, count := range colors {
		if count < minCount || (count == minCount && color < minColor) {
			minCount = count
			minColor = color
		}
	}

	return minColor
}
