package bodobia

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func bellmanFord(n, m, t int, hasSystem string, edges [][]int) int {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[1] = 0

	for i := 0; i < n-1; i++ {
		for _, edge := range edges {
			u, v, w := edge[0], edge[1], edge[2]
			if hasSystem[u-1] == '1' {
				dist[v] = min(dist[v], dist[u]+w)
			}
			if hasSystem[v-1] == '1' {
				dist[u] = min(dist[u], dist[v]-w)
			}
		}
	}

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if dist[u]+w < dist[v] && hasSystem[v-1] == '1' {
			return math.MinInt32
		}
	}
	return dist[n]
}

func main() {
	var n, m, t int
	fmt.Scan(&n, &m, &t)

	var hasSystem string
	fmt.Scan(&hasSystem)

	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		edges[i] = []int{u, v, w}
	}

	minTime := bellmanFord(n, m, t, hasSystem, edges)
	if minTime == math.MinInt32 {
		fmt.Println("Impossible")
	} else {
		fmt.Println(minTime)
	}
}
