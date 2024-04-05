package roozshomar

import "fmt"

func maxProfit(days int, profits []int) int {
	maxProfit := 0

	for i := 0; i < days; i++ {
		currentProfit := 0
		for j := i; j < days; j++ {
			currentProfit += profits[j]
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}

	return maxProfit
}

func main() {
	var days int
	fmt.Scanln(&days)

	profits := make([]int, days)
	for i := 0; i < days; i++ {
		fmt.Scan(&profits[i])
	}

	result := maxProfit(days, profits)
	fmt.Println(result)
}
