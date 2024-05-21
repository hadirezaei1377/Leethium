package pele

import "fmt"

func countWays(n int) int {

	ways := make([]int, n+1)

	ways[0] = 1
	ways[1] = 1

	for i := 2; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
		if i >= 5 {
			ways[i] += ways[i-5]
		}
	}

	return ways[n]
}

func main() {
	var n int
	fmt.Print("لطفاً تعداد پله‌ها را وارد کنید: ")
	fmt.Scan(&n)

	result := countWays(n)
	fmt.Println("تعداد راه‌های ممکن برای رسیدن به پله", n, ":", result)
}
