package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	times := make([][2]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&times[i][0], &times[i][1])
	}

	maxTime := 0
	for _, t := range times {
		if t[1] > maxTime {
			maxTime = t[1]
		}
	}

	totalCost := 0
	for t := 1; t <= maxTime; t++ {
		count := 0
		for _, time := range times {
			if t >= time[0] && t <= time[1] {
				count++
			}
		}
		if count == 1 {
			totalCost += a
		} else if count == 2 {
			totalCost += b
		} else if count == 3 {
			totalCost += c
		}
	}

	fmt.Println(totalCost)
}
