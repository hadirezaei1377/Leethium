package pooyan

import (
	"fmt"
	"sort"
)

type Customer struct {
	Preferences []Preference
}

type Preference struct {
	Size   int
	Filled bool
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	customers := make([]Customer, m)
	for i := 0; i < m; i++ {
		var k int
		fmt.Scan(&k)
		preferences := make([]Preference, k)
		for j := 0; j < k; j++ {
			var size, filled int
			fmt.Scan(&size, &filled)
			preferences[j] = Preference{Size: size, Filled: filled == 1}
		}
		customers[i] = Customer{Preferences: preferences}
	}

	result := findMinimumFilledTiles(n, customers)
	if result == nil {
		fmt.Println("IMPOSSIBLE")
	} else {
		for _, tile := range result {
			fmt.Printf("%d ", tile)
		}
	}
}

func findMinimumFilledTiles(n int, customers []Customer) []int {
	tiles := make([]int, n)

	for _, customer := range customers {
		for _, preference := range customer.Preferences {
			if !preference.Filled {
				tiles[preference.Size-1]++
			}
		}
	}

	result := make([]int, n)
	copy(result, tiles)

	sort.Ints(tiles)

	for i, customer := range customers {
		neededTiles := make(map[int]bool)
		for _, preference := range customer.Preferences {
			neededTiles[preference.Size-1] = true
		}

		for _, preference := range customer.Preferences {
			if !preference.Filled && tiles[preference.Size-1] == 1 && neededTiles[preference.Size-1] {
				result[preference.Size-1]--
				break
			}
		}

		if i == len(customers)-1 && len(tiles) > 0 && tiles[0] == 0 {
			return nil
		}
	}

	return result
}
